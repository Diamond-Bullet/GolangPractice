package design_pattern

/*
简单工厂模式

并非标准意义的工厂模式
基于接口，而不是基于共同的父类

Factory：工厂角色					负责实现创建所有实例的内部逻辑
Product：抽象产品角色				所创建的所有对象的父类，负责描述所有实例所共有的公共接口
ConcreteProduct：具体产品角色		创建目标，所有创建的对象都充当这个角色的某个具体类的实例。
*/

type Factory interface {
	Say(string) string
}

func NewFactory(fType string) Factory {
	switch fType {
	case "Car":
		return &CarFactory{Name: fType}
	case "Toy":
		return &ToyFactory{Name: fType}
	}
	panic("invalid type for factory")
}

type CarFactory struct {
	Name string
}

func (c *CarFactory) Say(s string) string {
	return "!!! I AM A " + c.Name
}

type ToyFactory struct {
	Name string
}

func (t *ToyFactory) Say(s string) string {
	return "!!! I AM A " + t.Name
}
