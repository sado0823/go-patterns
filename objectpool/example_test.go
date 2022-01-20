package objectpool

func ExampleNew() {
	pool := New(3)

	select {
	case obj := <-pool.ch:
		obj.Do()
	default:
		// 无可用对象
		return
	}
	// Output:
	// object-pool
}
