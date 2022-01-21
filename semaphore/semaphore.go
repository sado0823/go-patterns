package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrNoTicket          = errors.New("no ticket")
	ErrNoTicketToRelease = errors.New("no ticket to release")
)

type (
	Sema interface {
		Acquire(ctx context.Context) error
		Release(ctx context.Context) error
	}

	Worker struct {
		ch chan struct{}
	}
)

func New(num int) Sema {
	return &Worker{
		ch: make(chan struct{}, num),
	}
}

func (w *Worker) Acquire(ctx context.Context) error {
	for {
		select {
		case w.ch <- struct{}{}:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		default:
			continue
		}
	}

}

func (w *Worker) Release(ctx context.Context) error {
	for {
		select {
		case <-w.ch:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		default:
			continue
		}
	}
}

func main() {
	sema := New(1)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	err := sema.Acquire(ctx)
	fmt.Printf("err: %+v \n", err)

	err = sema.Acquire(ctx)
	fmt.Printf("err: %+v \n", err)
}
