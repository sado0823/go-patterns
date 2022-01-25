package main

import (
	"context"
	"fmt"
)

type (
	Chain struct {
		chain HandlerChain
		tmp   HandlerChain
	}

	HandlerChain interface {
		// setNext set next handler
		setNext(h HandlerChain) HandlerChain

		// runChain run all this chain handler
		runChain(ctx context.Context) error

		// current handler func
		handle(ctx context.Context) error
	}
)

func NewChain() *Chain {
	return &Chain{chain: &startHandler{}}
}

func (e *Chain) Next(h HandlerChain) *Chain {
	if e.tmp == nil {
		e.chain.setNext(h)
	} else {
		e.tmp.setNext(h)
	}
	e.tmp = h
	return e
}

func (e *Chain) Exec(ctx context.Context) error {
	return e.chain.runChain(ctx)
}

type baseHandler struct {
	next HandlerChain
}

func (b *baseHandler) setNext(h HandlerChain) HandlerChain {
	b.next = h
	return h
}

func (b *baseHandler) runChain(ctx context.Context) error {
	if b.next != nil {
		if err := b.next.handle(ctx); err != nil {
			return err
		}

		return b.next.runChain(ctx)
	}

	return nil
}

type startHandler struct {
	baseHandler
}

func (c *startHandler) handle(ctx context.Context) error {
	return nil
}

type CheckHandler struct {
	baseHandler
}

func (c *CheckHandler) handle(ctx context.Context) error {
	fmt.Println("do check")
	return nil
}

type RecheckHandler struct {
	baseHandler
}

func (c *RecheckHandler) handle(ctx context.Context) error {
	fmt.Println("do recheck")
	return nil
}

type SubmitHandler struct {
	baseHandler
}

func (c *SubmitHandler) handle(ctx context.Context) error {
	fmt.Println("do submit")
	return nil
}

func main() {
	err := NewChain().
		Next(&CheckHandler{}).
		Next(&SubmitHandler{}).
		Next(&RecheckHandler{}).
		Exec(context.Background())

	fmt.Println("err: ", err)
}
