package design_pattern

import "fmt"

/*
如果你的业务中出现了要依据不同的产品家族来生产其旗下的一系列产品的时候，抽象工厂模式就配上用场了。

例如小米公司和苹果公司就是两个不同产品家族，而他们两家都生产笔记本电脑和手机，
那么小米的笔记本电脑和苹果的笔记本电脑肯定不一样，手机情况也是如此。这就构成了两个产品家族的系列产品之间比较的关系。
*/

type Computer interface {
	SetCOS()
}

type MacComputer struct{}

func (m *MacComputer) SetCOS() {
	fmt.Println("MacComputer: OX")
}

type MiComputer struct{}

func (m *MiComputer) SetCOS() {
	fmt.Println("MiComputer: Windows")
}

type Phone interface {
	SetPOS()
}

type MacPhone struct{}

func (m *MacPhone) SetPOS() {
	fmt.Println("MacPhone: apple")
}

type MiPhone struct{}

func (m *MiPhone) SetPOS() {
	fmt.Println("MiPhone: Android")
}

type AbstractFactory interface {
	MakeComputer() Computer
	MakePhone() Phone
}

type MacFactory struct{}

func (m *MacFactory) MakeComputer() Computer {
	return &MacComputer{}
}

func (m *MacFactory) MakePhone() Phone {
	return &MacPhone{}
}

// TODO if you want to implement MiFactory
