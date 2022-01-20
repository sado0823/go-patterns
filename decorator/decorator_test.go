package decorator

import (
	"testing"
)

func TestCalculate(t *testing.T) {

	sq := Calculate(Sqrt)
	do := Calculate(Double)
	p5 := Calculate(Plus5)

	v := p5(do(sq(100))) // sq100 * 2 + 5
	if v != 25 {
		t.Fatalf("expected:%d get:%v", 25, v)
	}
}
