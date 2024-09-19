package structural_pattern

import (
	"strconv"
	"testing"
)

/*
Wrapper is just like its name, joining two incompatible interfaces together.

It has an alias `Adapter`, namely adding an extra layer between two interfaces.

There is a saying in IT world: there is no problem that cannot be solved by adding another layer.
*/

func TestWrapper(t *testing.T) {
	wrappee := NewWrappee()

	wrapper := NewWrapper(wrappee)
	println(wrapper.Request(1))
}

// Target
type Target interface {
	Request(request int) string
}

// Wrapper
type Wrapper struct {
	Wrappee
}

// Request encapsulation and decoupling
func (a *Wrapper) Request(request int) string {
	legacyRequest := strconv.Itoa(request)
	return a.LegacyRequest(legacyRequest)
}

func NewWrapper(wrappee Wrappee) *Wrapper {
	return &Wrapper{
		wrappee,
	}
}

// Wrappee
type Wrappee interface {
	LegacyRequest(request string) string
}

// NewWrappee
func NewWrappee() Wrappee {
	return &Wrappee1{}
}

// Wrappee1
type Wrappee1 struct{}

func (a *Wrappee1) LegacyRequest(request string) string {
	return request
}
