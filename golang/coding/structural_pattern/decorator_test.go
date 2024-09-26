package structural_pattern

import (
	"GolangPractice/utils/logger"
	"github.com/gookit/color"
	"testing"
)

// Decorator
// There are two means to change the behavior of an object: inheritance and composition.
// Go doesn't support inheritance. So we use composition and overwrite the method of embedded struct, to change the behavior of an object.
// That is what Decorator is.

func TestDecorator(t *testing.T) {
	adder := &SimpleAdder{}
	adder.Add(1, 2)

	printAdder := &PrintDecorator{Adder: adder}
	printAdder.Add(1, 2)

	logAdder := &LogDecorator{Adder: printAdder}
	logAdder.Add(1, 2)
}

type Add interface {
	Add(a, b int) int
}

type SimpleAdder struct{}

func (s *SimpleAdder) Add(a, b int) int {
	return a + b
}

type PrintDecorator struct {
	Adder Add
}

func (p *PrintDecorator) Add(a, b int) int {
	res := p.Adder.Add(a, b)
	color.Redln("a + b =", res)
	return res
}

type LogDecorator struct {
	Adder Add
}

func (l *LogDecorator) Add(a, b int) int {
	logger.Infoln("a =", a, "b =", b)
	res := l.Adder.Add(a, b)
	return res
}