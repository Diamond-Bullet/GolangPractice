package structural_pattern

import (
	"GolangPractice/utils/logger"
	"strconv"
	"testing"
)

/*
Adapter is just like its name, joining two incompatible interfaces together.
It has an alias `Wrapper`, namely adding an extra layer between two interfaces.

There is a saying in IT world: there is no problem that cannot be solved by adding another layer.
*/

func TestAdapter(t *testing.T) {
	adapter := NewAdapter(NewAdaptee())
	logger.Info(adapter.Request(1))
}

type Target interface {
	Request(request int) string
}

type Adaptee interface {
	LegacyRequest(request string) string
}

type Adapter struct {
	Adaptee
}

// Request encapsulation and decoupling
func (a *Adapter) Request(request int) string {
	legacyRequest := strconv.Itoa(request)
	return a.LegacyRequest(legacyRequest)
}

func NewAdapter(adaptee Adaptee) *Adapter {
	return &Adapter{
		adaptee,
	}
}

func NewAdaptee() Adaptee {
	return &LegacyAdaptee{}
}

type LegacyAdaptee struct{}

func (a *LegacyAdaptee) LegacyRequest(request string) string {
	return request
}
