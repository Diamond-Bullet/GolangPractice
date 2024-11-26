package creational_pattern

import (
	"GolangPractice/lib/logger"
	"fmt"
	"testing"
)

/*
Builder:
	The constructor of a class has both compulsory and optional parameters, and typically more than 4.
	The use and function are similar to `golang/coding/option_test.go`.
	You can customize the building of a class rather than filling in all the parameters.

	It resembles the process of building a car, in which you build all components separately and then assemble them together.

Roles:
	1. Product: The final object that will be created by the builder.
	2. Builder: The interface that will define the steps to build the product.
	3. ConcreteBuilder: The class that implements the Builder interface.
	4. Director: The class that will construct the final object using the Builder interface.
*/

func TestBuilder(t *testing.T) {
	macBuilder := NewMacBuilder("Intel Core i7-10900", "Santax 16GB DDR4")
	director := NewDirector(macBuilder)

	mac := director.Construct("MAC Standard KeyBoard", "SkyWorth 15.6", 6)
	logger.Info(mac.Info())
}

type Mac struct {
	CPU, RAM          string
	USBCount          int
	KeyBoard, Display string
}

func (m *Mac) Info() string {
	return fmt.Sprintf(`
	CPU:		%s
	RAM:		%s
	USBCOUNT:	%d
	KEYBOARD:	%s
	DISPLAY:	%s
`, m.CPU, m.RAM, m.USBCount, m.KeyBoard, m.Display)
}

// Builder
type Builder interface {
	SetCPU(cpu string) Builder
	SetRAM(ram string) Builder
	WithUsbCount(int) Builder
	WithKeyBoard(string) Builder
	WithDisplay(string) Builder
	Build() *Mac
}

type MacBuilder struct {
	Mac *Mac
}

func NewMacBuilder(cpu, ram string) Builder {
	return (&MacBuilder{Mac: &Mac{}}).SetCPU(cpu).SetRAM(ram)
}

func (m *MacBuilder) SetCPU(cpu string) Builder {
	m.Mac.CPU = cpu
	return m
}

func (m *MacBuilder) SetRAM(ram string) Builder {
	m.Mac.RAM = ram
	return m
}

func (m *MacBuilder) WithUsbCount(usbCount int) Builder {
	m.Mac.USBCount = usbCount
	return m
}

func (m *MacBuilder) WithKeyBoard(keyBoard string) Builder {
	m.Mac.KeyBoard = keyBoard
	return m
}

func (m *MacBuilder) WithDisplay(display string) Builder {
	m.Mac.Display = display
	return m
}

func (m *MacBuilder) Build() *Mac {
	return m.Mac
}

// Director
type Director struct {
	builder Builder
}

func (d *Director) Construct(keyBoard, DisPlay string, UsbCount int) *Mac {
	d.builder.WithKeyBoard(keyBoard).WithDisplay(DisPlay).WithUsbCount(UsbCount)
	return d.builder.Build()
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}
