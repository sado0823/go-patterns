package singleton

import (
	"testing"
)

func Test_Hungry(t *testing.T) {

	get := GetHungrySingleton()
	if get.Get() != 1 {
		t.Fatalf("expected:%d, get:%d", 1, get.Get())
	}

	get.Set(2)

	get2 := GetHungrySingleton()
	if get2.Get() != 2 {
		t.Fatalf("expected:%d, get:%d", 2, get2.Get())
	}
}
