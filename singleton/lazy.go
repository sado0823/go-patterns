package singleton

import "sync"

// Lazy 单例模式 - 懒汉模式
type Lazy struct {
	val int64
}

var (
	lazySingleton *Lazy
	once          = &sync.Once{}
)

func GetLazySingleton() *Lazy {
	once.Do(func() {
		lazySingleton = &Lazy{val: 1}
	})

	return lazySingleton
}

func (h *Lazy) Get() int64 {
	return h.val
}

func (h *Lazy) Set(v int64) {
	h.val = v
}
