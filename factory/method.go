package factory

const (
	TypeA MethodType = iota
	TypeB
	TypeC
)

// Method 工厂方法模式
type (
	MethodType int64

	Method struct{}

	MethodEvent interface {
		Name() string
	}

	MethodEventA struct{}
	MethodEventB struct{}
	MethodEventC struct{}
)

// Create 实现方式一: 一个create方法创建所有的类型
func (m *Method) Create(eType MethodType) MethodEvent {
	switch eType {
	case TypeA:
		return &MethodEventA{}
	case TypeB:
		return &MethodEventB{}
	case TypeC:
		return &MethodEventC{}
	default:
		return &MethodEventA{}
	}
}

// CreateA 实现方式二: 为每个类型提供一个创建方法
func (m *Method) CreateA() MethodEvent {
	return &MethodEventA{}
}

func (m *Method) CreateB() MethodEvent {
	return &MethodEventB{}
}

func (m *Method) CreateC() MethodEvent {
	return &MethodEventC{}
}

func (m *MethodEventA) Name() string {
	return "MethodEventA"
}

func (m *MethodEventB) Name() string {
	return "MethodEventB"
}

func (m *MethodEventC) Name() string {
	return "MethodEventC"
}
