package builder

import "testing"

func TestNewBuilder(t *testing.T) {
	msg := NewMsgBuilder().
		WithKV("foo", "bar").
		WithSpan("span").
		WithTrace("trace").
		Build()

	if msg.kv["foo"] != "bar" {
		t.Fatalf("expected:%s, get:%s", "bar", msg.kv["foo"])
	}

	if msg.meta.span != "span" {
		t.Fatalf("expected:%s, get:%s", "span", msg.meta.span)
	}

	if msg.meta.trace != "trace" {
		t.Fatalf("expected:%s, get:%s", "trace", msg.meta.trace)
	}
}
