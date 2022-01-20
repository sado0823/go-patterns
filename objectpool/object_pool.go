package objectpool

import "fmt"

type (
	Obj  struct{}
	Pool struct {
		size int
		ch   chan *Obj
	}
)

func (o *Obj) Do() {
	fmt.Println("object-pool")
}

func New(size int) *Pool {
	p := &Pool{
		size: size,
		ch:   make(chan *Obj, size),
	}

	for i := 0; i < size; i++ {
		p.ch <- new(Obj)
	}

	return p
}

func (p *Pool) Get() *Obj {
	return <-p.ch
}
