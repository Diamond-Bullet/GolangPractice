package design_pattern

/*
将一个接口转换为客户端所期待的接口，从而使两个接口不兼容的类可以在一起工作

适配器模式还有个别名叫：Wrapper（包装器），顾名思义就是将目标类用一个新类包装一下，相当于在客户端与目标类直接加了一层。

IT世界有句俗语：没有什么问题是加一层不能解决的。
*/

// Target 目标接口
type Target interface {
	Request() string
}

// Adaptee 需被适配的接口
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee 需被适配的接口的 工厂
func NewAdaptee() Adaptee {
	return &Adaptee1{}
}

// Adaptee1 具体的需适配接口
type Adaptee1 struct{}

func (a *Adaptee1) SpecificRequest() string {
	return "Adaptee1 method"
}

// Adaptor 适配器
type Adaptor struct {
	Adaptee
}

// Request 包了一下方法，解耦
func (a *Adaptor) Request() string {
	return a.SpecificRequest()
}

func NewAdaptor(adaptee Adaptee) *Adaptor {
	return &Adaptor{
		adaptee,
	}
}
