package design_pattern

import (
	"GolangPractice/utils/logger"
	"testing"
)

/*
Simple Factory: not a standard design pattern but an approach to centralize the creation of objects.
*/

func TestFactory(t *testing.T) {
	car := NewProduct(Car)
	logger.Infoln(car.Say())

	toy := NewProduct(Toy)
	logger.Infoln(toy.Say())
}

type Product interface {
	Say() string
}

func NewProduct(productType ProductType) Product {
	switch productType {
	case Car:
		return &CAR{Name: productType.String()}
	case Toy:
		return &TOY{Name: productType.String()}
	}
	return nil
}

type CAR struct {
	Name string
}

func (c *CAR) Say() string {
	return "!!! I AM A " + c.Name
}

type TOY struct {
	Name string
}

func (t *TOY) Say() string {
	return "!!! I AM A " + t.Name
}

type ProductType string

func (p ProductType) String() string {
	return string(p)
}

const (
	Car ProductType = "Car"
	Toy             = "Toy"
)