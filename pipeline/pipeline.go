package main

import (
	"fmt"
	"go-version/src/math/rand"
)

func generate(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)

		for {
			for _, value := range values {
				select {
				case <-done:
					return
				case valueStream <- value:

				}
			}
		}
	}()
	return valueStream
}

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})

	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
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

func toString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)

	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			default:
				s, ok := v.(string)
				if !ok {
					continue
				}
				select {
				case <-done:
					return
				case stringStream <- s:
				}
			}
		}
	}()
	return stringStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	// a:
	// take:  1
	// take:  2
	// take:  3
	takeStream := take(done, generate(done, 1, 2, 3, 4, 5), 3)
	for t := range takeStream {
		fmt.Println("take: ", t)
	}

	// b:
	// rand int:  5577006791947779410
	// rand int:  8674665223082153551
	// rand int:  6129484611666145821
	// rand int:  4037200794235010051
	// rand int:  3916589616287113937
	stream := take(done, repeatFn(done, func() interface{} {
		return rand.Int()
	}), 5)
	for randI := range stream {
		fmt.Println("rand int: ", randI)
	}

	// c
	// str:  a
	// str:  b
	// str:  a
	// str:  b
	// str:  a
	strStream := toString(done, take(done, repeat(done, "a", "b"), 5))
	for str := range strStream {
		fmt.Println("str: ", str)
	}
}
