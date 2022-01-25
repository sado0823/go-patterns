package main

import (
	"fmt"
	"time"
)

const (
	File int64 = iota
	Std
	Db
)

type (
	Event interface {
		Do() int64
	}

	// Observer 观察者模式
	Observer interface {
		OnNotify(e Event)
	}

	Notifier interface {
		Register(ob Observer)
		Deregister(ob Observer)
		Notify(e Event)
	}

	FileEvent struct{}
	StdEvent  struct{}
	DbEvent   struct{}

	EventObserver struct {
		fn func(e Event)
	}

	OutPutNotifier struct {
		store map[Observer]struct{}
	}
)

func NewOutPutNotifier() *OutPutNotifier {
	return &OutPutNotifier{store: map[Observer]struct{}{}}
}

func NewEventObserver(fn func(e Event)) *EventObserver {
	return &EventObserver{fn}
}

func (o *OutPutNotifier) Register(ob Observer) {
	o.store[ob] = struct{}{}
}

func (o *OutPutNotifier) Deregister(ob Observer) {
	delete(o.store, ob)
}

func (o *OutPutNotifier) Notify(e Event) {
	for observer := range o.store {
		observer.OnNotify(e)
	}
}

func (f *EventObserver) OnNotify(e Event) {
	f.fn(e)
}

func (f *FileEvent) Do() int64 {
	return File
}

func (f *StdEvent) Do() int64 {
	return Std
}

func (f *DbEvent) Do() int64 {
	return Db
}

func main() {
	notifier := NewOutPutNotifier()
	notifier.Register(&EventObserver{fn: func(e Event) {
		fmt.Printf("do something A: 🛩%d号 \n", e.Do())
	}})
	notifier.Register(&EventObserver{func(e Event) {
		fmt.Printf("do something B: 🏡%d号 \n", e.Do())
	}})

	timer15s := time.NewTimer(time.Second * 6)
	ticker1s := time.NewTicker(time.Second * 1)
	ticker5s := time.NewTicker(time.Second * 2)
	defer func() {
		timer15s.Stop()
		ticker1s.Stop()
		ticker5s.Stop()
	}()

	for {
		select {
		case <-timer15s.C:
			notifier.Notify(&StdEvent{})
			return
		case <-ticker1s.C:
			notifier.Notify(&FileEvent{})
		case <-ticker5s.C:
			notifier.Notify(&DbEvent{})
		default:
			continue
		}
	}
}
