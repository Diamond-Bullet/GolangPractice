package behavioral_pattern

import "testing"

// Visitor provide a way to add operations to an object while not changing it.

func TestVisitor(t *testing.T) {
	info := &InfoHost{
		F1: "f1",
		F2: "f2",
		F3: "f3",
	}
	f1Visitor := &F1Visitor{}
	info.Accept(f1Visitor)
}

type Host interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(host Host)
}

type InfoHost struct {
	F1 string
	F2 string
	F3 string
}

func (i *InfoHost) Accept(visitor Visitor) {
	visitor.Visit(i)
}

type F1Visitor struct{}

func (f *F1Visitor) Visit(host Host) {

}
