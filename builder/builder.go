package builder

import "sync"

// Builder 建造者模式
type Builder struct {
	once *sync.Once
	msg  *Msg
}

type Msg struct {
	kv   map[string]string
	meta *struct {
		trace string
		span  string
	}
}

func NewMsgBuilder() *Builder {
	return &Builder{
		once: &sync.Once{},
		msg: &Msg{
			kv: nil,
			meta: &struct {
				trace string
				span  string
			}{trace: "", span: ""},
		},
	}
}

func (b Builder) Build() *Msg {
	return b.msg
}

func (b *Builder) WithKV(key string, val string) *Builder {
	b.once.Do(func() {
		b.msg.kv = make(map[string]string)
	})

	b.msg.kv[key] = val
	return b
}

func (b *Builder) WithTrace(trace string) *Builder {
	b.msg.meta.trace = trace

	return b
}

func (b *Builder) WithSpan(span string) *Builder {
	b.msg.meta.span = span

	return b
}
