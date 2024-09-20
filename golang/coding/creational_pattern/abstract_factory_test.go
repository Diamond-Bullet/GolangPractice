package creational_pattern

import (
	"fmt"
	"testing"
)

/*
Abstract Factory:
A factory implementation is responsible for the creation of an object family.
Here the object family is a set of objects with correlations.

Roles:
	1. Abstract Factory: The interface of the factory.
	2. Concrete Factory: The implementation of the factory.
	3. Abstract Product: The interface of the product.
	4. Concrete Product: The implementation of the product.
*/

func TestAbstractFactory(t *testing.T) {
	appleFactory := &AppleFactory{}
	macComputer := appleFactory.MakeComputer()
	macPhone := appleFactory.MakePhone()
	macComputer.SetCOS()
	macPhone.SetPOS()

	miFactory := &MiFactory{}
	miComputer := miFactory.MakeComputer()
	miPhone := miFactory.MakePhone()
	miComputer.SetCOS()
	miPhone.SetPOS()
}

type AbstractFactory interface {
	MakeComputer() Computer
	MakePhone() Phone
}

type Computer interface {
	SetCOS()
}

type Phone interface {
	SetPOS()
}

type AppleFactory struct{}

func (m *AppleFactory) MakeComputer() Computer {
	return &MacComputer{}
}

func (m *AppleFactory) MakePhone() Phone {
	return &IPhone{}
}

type MacComputer struct{}

func (m *MacComputer) SetCOS() {
	fmt.Println("MacComputer: OX")
}

type IPhone struct{}

func (m *IPhone) SetPOS() {
	fmt.Println("IPhone: apple")
}

type MiFactory struct{}

func (m *MiFactory) MakeComputer() Computer {
	return &MiComputer{}
}

func (m *MiFactory) MakePhone() Phone {
	return &MiPhone{}
}

type MiComputer struct{}

func (m *MiComputer) SetCOS() {
	fmt.Println("MiComputer: Windows")
}

type MiPhone struct{}

func (m *MiPhone) SetPOS() {
	fmt.Println("MiPhone: Android")
}
