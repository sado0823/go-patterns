package decorator

import (
	"fmt"
	"math"
)

type Decorator func(i float64) float64

func Calculate(de Decorator) Decorator {
	return func(i float64) float64 {
		fmt.Println("start")
		v := de(i)
		fmt.Println("end")
		return v
	}
}

func Sqrt(i float64) float64 {
	fmt.Println("do sqrt")
	return math.Sqrt(i)
}

func Double(i float64) float64 {
	fmt.Println("do double")
	return i * 2
}

func Plus5(i float64) float64 {
	fmt.Println("do plus5")
	return i + 5
}
