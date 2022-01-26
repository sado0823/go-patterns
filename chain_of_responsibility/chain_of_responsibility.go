package main

import (
	"context"
	"fmt"
)

type (
	Chain struct {
		chain Chainer
		tmp   Chainer
	}

	Chainer interface {
		// setNext set next handler
		setNext(h Chainer) Chainer

		// runChain run all this chain handler
		runChain(ctx context.Context) error

		// current handler func
		handle(ctx context.Context) error
	}
)

func NewChain() *Chain {
	return &Chain{chain: &startChainer{}}
}

func (e *Chain) Next(h Chainer) *Chain {
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

type baseChainer struct {
	next Chainer
}

func (b *baseChainer) setNext(h Chainer) Chainer {
	b.next = h
	return h
}

func (b *baseChainer) runChain(ctx context.Context) error {
	if b.next != nil {
		if err := b.next.handle(ctx); err != nil {
			return err
		}

		return b.next.runChain(ctx)
	}

	return nil
}

type startChainer struct {
	baseChainer
}

func (c *startChainer) handle(ctx context.Context) error {
	return nil
}

type CheckChainer struct {
	baseChainer
}

func (c *CheckChainer) handle(ctx context.Context) error {
	fmt.Println("do check")
	return nil
}

type RecheckChainer struct {
	baseChainer
}

func (c *RecheckChainer) handle(ctx context.Context) error {
	fmt.Println("do recheck")
	return nil
}

type SubmitChainer struct {
	baseChainer
}

func (c *SubmitChainer) handle(ctx context.Context) error {
	fmt.Println("do submit")
	return nil
}

func main() {
	err := NewChain().
		Next(&CheckChainer{}).
		Next(&SubmitChainer{}).
		Next(&RecheckChainer{}).
		Exec(context.Background())

	fmt.Println("err: ", err)
}
