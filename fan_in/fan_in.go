package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var (
		inStream = make(chan interface{})
		wg       sync.WaitGroup
	)

	inFn := func(c <-chan interface{}) {
		defer wg.Done()
		for v := range c {
			select {
			case <-done:
				return
			case inStream <- v:
			}
		}
	}

	wg.Add(len(channels))
	for _, channel := range channels {
		go inFn(channel)
	}

	go func() {
		wg.Wait()
		close(inStream)
	}()

	return inStream
}

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})

	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			default:
				s, ok := v.(int)
				if !ok {
					continue
				}
				select {
				case <-done:
					return
				case intStream <- s:
				}
			}
		}
	}()
	return intStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, count int) <-chan interface{} {
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)

		for i := 0; i < count; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func primeFinder(done <-chan interface{}, intChan <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})

	primeFn := func(v int) bool {
		for i := 2; i < v-1; i++ {
			if v%i == 0 {
				return false
			}
		}
		return true
	}

	go func() {
		defer close(primeStream)
		for v := range intChan {
			select {
			case <-done:
				return
			default:
				if primeFn(v) {
					select {
					case <-done:
						return
					case primeStream <- v:
					}
				}
			}

		}
	}()

	return primeStream
}

func main() {
	noFanIn := func() time.Duration {
		randFn := func() interface{} { return rand.Intn(50000000) }

		done := make(chan interface{})
		defer close(done)

		start := time.Now()
		randIntStream := toInt(done, repeatFn(done, randFn))

		// no fan-in
		fmt.Println("Primes:")
		for prime := range take(done, primeFinder(done, randIntStream), 100) {
			fmt.Printf("\t%d\n", prime)
		}

		return time.Since(start)
	}

	withFanIn := func() time.Duration {
		randFn := func() interface{} { return rand.Intn(50000000) }

		done := make(chan interface{})
		defer close(done)

		start := time.Now()
		randIntStream := toInt(done, repeatFn(done, randFn))

		// with fan-in
		numFinders := runtime.NumCPU()
		fmt.Printf("Spinning up %d prime finders.\n", numFinders)
		finders := make([]<-chan interface{}, numFinders)
		fmt.Println("Primes:")
		for i := 0; i < numFinders; i++ {
			finders[i] = primeFinder(done, randIntStream)
		}

		for prime := range take(done, fanIn(done, finders...), 100) {
			fmt.Printf("\t%d\n", prime)
		}

		return time.Since(start)
	}

	var (
		wg       sync.WaitGroup
		no, with time.Duration
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		no = noFanIn()
	}()
	go func() {
		defer wg.Done()
		with = withFanIn()
	}()
	wg.Wait()
	fmt.Println("no fan-in cost: ", no)
	fmt.Println("with fan-in cost: ", with)

}
