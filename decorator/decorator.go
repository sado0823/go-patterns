package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"
)

type Decorator func(i float64) float64

func Calculate(i float64) float64 {
	fmt.Println("Calculate start")
	defer fmt.Println("Calculate end")
	return math.Sqrt(i)
}

func wrapLogger(de Decorator, logger *log.Logger) Decorator {
	return func(i float64) float64 {
		fmt.Println("logger start")
		defer func(t time.Time) {
			logger.Printf("do logger - end: from:%v, to:%v, diff:%v  \n", t, time.Now(), time.Since(t).Microseconds())
			fmt.Println("logger end")
		}(time.Now())

		return de(i)
	}
}

func wrapCache(de Decorator, cache *sync.Map) Decorator {
	return func(i float64) float64 {
		fmt.Println("cache start")
		defer fmt.Println("cache end")

		key := fmt.Sprintf("%v", i)

		if v, ok := cache.Load(key); ok {
			fmt.Println("hit cache")
			return v.(float64)
		}

		res := de(i)
		cache.Store(key, res)
		return res
	}
}

func main() {
	l := wrapLogger(Calculate, log.New(os.Stdout, "test logger stdout ", 1))
	c := wrapCache(l, &sync.Map{})

	c(25)
	fmt.Println("==================")
	c(25)
	fmt.Println("==================")
	c(36)

}
