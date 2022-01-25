package main

import (
	"context"
	"errors"
	"math"
	"testing"
	"time"
)

func TestWorker_Acquire_ONE(t *testing.T) {
	sema := New(1)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*2000)
	defer cancelFunc()

	err := sema.Acquire(ctx)
	if err != nil {
		t.Fatalf("expected err nil, got:%v", err)
	}
	defer func() {
		_ = sema.Release(context.Background())
	}()

	err = sema.Acquire(ctx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected err context deadline, got err: %v", err)
	}

}

func TestWorker_Acquire_MORE(t *testing.T) {
	sema := New(2)
	err := sema.Acquire(context.Background())
	if err != nil {
		t.Fatalf("expected err nil, got:%v", err)
	}

	err = sema.Acquire(context.Background())
	if err != nil {
		t.Fatalf("expected err nil, got:%v", err)
	}

	go func() {
		time.Sleep(time.Millisecond * 50)
		_ = sema.Release(context.Background())
	}()

	err = sema.Acquire(context.Background())
	if err != nil {
		t.Fatalf("expected err nil, got:%v", err)
	}
}

func TestWorker_Try(t *testing.T) {
	sema := New(1)
	err := sema.Acquire(context.Background())
	if err != nil {
		t.Fatalf("expected err nil, got:%v", err)
	}

	now := time.Now()
	go func() {
		time.Sleep(time.Second * 2)
		_ = sema.Release(context.Background())
	}()

	err = sema.Acquire(context.Background())
	diff := time.Since(now).Seconds()
	if err != nil {
		t.Fatalf("expected err nil, got:%v", err)
	}

	if v := math.Round(diff); v != 2 {
		t.Fatalf("try should wait 500ms, but get:%v", v)
	}

}

func TestWorker_Release(t *testing.T) {
	sema := New(2)
	err := sema.Acquire(context.Background())
	if err != nil {
		t.Fatalf("expected err nil, got:%v", err)
	}

	err = sema.Release(context.Background())
	if err != nil {
		t.Fatalf("expected err nil, got:%v", err)
	}

	err = sema.Acquire(context.Background())
	if err != nil {
		t.Fatalf("expected err nil, got:%v", err)
	}

	err = sema.Release(context.Background())
	if err != nil {
		t.Fatalf("expected err nil, got:%v", err)
	}
}