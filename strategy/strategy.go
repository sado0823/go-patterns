package main

import (
	"fmt"
)

type (
	// Math 策略模式
	Math interface {
		Do(i, j int) int
	}

	Operator struct {
		m Math
	}

	Plus           struct{}
	Minus          struct{}
	Multiplication struct{}
)

func NewOperator(m Math) *Operator {
	return &Operator{m}
}

func (o *Operator) Calculate(i, j int) int {
	return o.m.Do(i, j)
}

func (p *Plus) Do(i, j int) int {
	return i + j
}

func (p *Minus) Do(i, j int) int {
	return i - j
}

func (p *Multiplication) Do(i, j int) int {
	return i * j
}

func main() {
	i, j := 10, 20
	v := NewOperator(&Plus{}).Calculate(i, j) // 30
	fmt.Println("res: ", v)

	v = NewOperator(&Minus{}).Calculate(i, j) // -10
	fmt.Println("res: ", v)

	v = NewOperator(&Multiplication{}).Calculate(i, j) // 200
	fmt.Println("res: ", v)

}
