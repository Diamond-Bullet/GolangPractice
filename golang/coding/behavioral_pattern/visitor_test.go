package behavioral_pattern

import (
	"github.com/gookit/color"
	"testing"
)

// Visitor provide a way to add operations to an object while not changing it.

func TestVisitor(t *testing.T) {
	info := &InfoHost{
		Name:    "INFO",
		Region:  "JAPAN",
		Replica: "3",
	}

	nameVisitor := &NameVisitor{}
	info.Accept(nameVisitor)
}

type Host interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(host Host)
}

type InfoHost struct {
	Name    string
	Region  string
	Replica string
}

func (i *InfoHost) Accept(visitor Visitor) {
	visitor.Visit(i)
}

type NameVisitor struct{}

func (n *NameVisitor) Visit(host Host) {
	switch h := host.(type) {
	case *InfoHost:
		if h == nil {
			return
		}
		color.Blueln(h.Name)
	default:
		return
	}
}
