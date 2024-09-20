package creational_pattern

import "testing"

/*
Factory Method
A dedicated factory implementation is applied to create A particular object, rather than all in ONE single factory class.
Now if we are going to add a new object type, we just need to create a new factory class for it instead of altering the existing one.
This feature makes the pattern more in line with the "Open-Closed Principle".

Roles:
	Product: interface of the product
	ConcreteProduct: implementation of the product
	Factory: interface of the factory
	ConcreteFactory: implementation of the factory
*/

func TestFactoryMethod(t *testing.T) {
	plusFactory := PlusOperatorFactory{}
	plus := plusFactory.Create()
	plus.SetA(1)
	plus.SetB(2)
	println(plus.Result())

	minusFactory := MinusOperatorFactory{}
	minus := minusFactory.Create()
	minus.SetA(1)
	minus.SetB(2)
	println(minus.Result())
}

// Operator encapsulated object
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory interface for creating an operator
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

// PlusOperatorFactory
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

// MinusOperatorFactory
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
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
