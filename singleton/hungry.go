package singleton

// Hungry 单例模式 - 饿汉模式
type Hungry struct {
	val int64
}

var hungrySingleton = &Hungry{val: 1}

func GetHungrySingleton() *Hungry {
	return hungrySingleton
}

func (h *Hungry) Get() int64 {
	return h.val
}

func (h *Hungry) Set(v int64) {
	h.val = v
}
