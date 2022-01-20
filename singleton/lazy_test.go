package singleton

import "testing"

func Test_Lazy(t *testing.T) {

	get := GetLazySingleton()
	if get.Get() != 1 {
		t.Fatalf("expected:%d, get:%d", 1, get.Get())
	}

	get.Set(2)

	get2 := GetLazySingleton()
	if get2.Get() != 2 {
		t.Fatalf("expected:%d, get:%d", 2, get2.Get())
	}
}
