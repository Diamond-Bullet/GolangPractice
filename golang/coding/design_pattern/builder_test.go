package design_pattern

import (
	"fmt"
	"testing"
)

/*
Builder
将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

复杂对象相当于一辆有待建造的汽车，而对象的属性相当于汽车的部件，建造产品的过程就相当于组合部件的过程。
由于组合部件的过程很复杂，因此，这些部件的组合过程往往被“外部化”到一个称作建造者的对象里，
建造者返还给客户端的是一个已经建造完毕的完整产品对象，而用户无须关心该对象所包含的属性以及它们的组装方式，这就是建造者模式的模式动机。

Use: The constructor of a class has both compulsory and optional parameters, and typically more than 4.
*/

func TestBuilder(t *testing.T) {
	mac := NewMac("Intel Core i7-10900", "Santax 16GB DDR4")
	director := NewDirector(mac)
	director.Construct("MAC Standard KeyBoard", "SkyWorth 15.6", 6)
	println(mac.GetComputer())
}

// Builder 生成器接口
type Builder interface {
	PartUsbCount(int)
	PartKeyBoard(string)
	PartDisplay(string)
}

type MacBuilder struct {
	CPU, RAM          string
	USBCount          int
	KeyBoard, Display string
}

func NewMac(cpu, ram string) *MacBuilder {
	return &MacBuilder{
		CPU: cpu,
		RAM: ram,
	}
}

func (m *MacBuilder) PartUsbCount(usbCount int) {
	m.USBCount = usbCount
}

func (m *MacBuilder) PartKeyBoard(keyBoard string) {
	m.KeyBoard = keyBoard
}

func (m *MacBuilder) PartDisplay(display string) {
	m.Display = display
}

func (m MacBuilder) GetComputer() string {
	return fmt.Sprintf(`
	CPU:		%s
	RAM:		%s
	USBCOUNT:	%d
	KEYBOARD:	%s
	DISPLAY:	%s
`, m.CPU, m.RAM, m.USBCount, m.KeyBoard, m.Display)
}

// Director 指导器，负责组装各部分
type Director struct {
	builder Builder
}

func (d *Director) Construct(keyBoard, DisPlay string, UsbCount int) {
	d.builder.PartKeyBoard(keyBoard)
	d.builder.PartDisplay(DisPlay)
	d.builder.PartUsbCount(UsbCount)
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}
