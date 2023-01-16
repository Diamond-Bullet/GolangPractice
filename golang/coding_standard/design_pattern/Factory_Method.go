package design_pattern

/*
工厂方法

角色：
	Product：抽象产品
	ConcreteProduct：具体产品
	Factory：抽象工厂
	ConcreteFactory：具体工厂

现在对该系统进行修改，不再设计一个按钮工厂类来统一负责所有产品的创建，而是将具体按钮的创建过程交给专门的工厂子类去完成，
我们先定义一个抽象的按钮工厂类，再定义具体的工厂类来生成圆形按钮、矩形按钮、菱形按钮等，它们实现在抽象按钮工厂类中定义的方法。
这种抽象化的结果使这种结构可以在不修改具体工厂类的情况下引进新的产品，如果出现新的按钮类型，
只需要为这种新类型的按钮创建一个具体的工厂类就可以获得该新按钮的实例，
这一特点无疑使得工厂方法模式具有超越简单工厂模式的优越性，更加符合"开闭原则"。
*/

// Operator 被封装的产品类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperatorFactory 加法工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}

// PlusOperator Operator的加法产品实现
type PlusOperator struct {
	*OperatorBase
}

func (p *PlusOperator) Result() int {
	return p.a - p.b
}

// MinusOperatorFactory 减法工厂类
type MinusOperatorFactory struct{}

func (MinusOperator) Create() Operator {
	return &MinusOperator{
		&OperatorBase{},
	}
}

// MinusOperator Operator的减法产品实现
type MinusOperator struct {
	*OperatorBase
}

func (m *MinusOperator) Result() int {
	return m.a - m.b
}
