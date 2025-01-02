package behavioral_pattern

import (
	"GolangPractice/pkg/logger"
	"testing"
)

/*
ProtoType

When the cost of building an object is rather high,
for example, we need to access the database to get the data,
but we have to build it multiple times.

If the object built all have the same state at start, We can build a prototype object with the state, and then adjust it.

In gorm, when we create a new transaction instance, we use Clone method.
*/

func TestPrototype(t *testing.T) {
	manager := NewPrototypeManager()

	t1 := &TypeA{name: "type1"}
	manager.Set("t1", t1)

	t2 := manager.Get("t1").Clone()
	if t1 == t2 {
		logger.Error("t1 == t2")
	}

	t3 := t1.Clone()
	if t1 == t3 {
		logger.Error("t1 == t3")
	}
}

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes: make(map[string]Cloneable)}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

type TypeA struct {
	name string
}

func (t *TypeA) Clone() Cloneable {
	tc := *t
	return &tc
}
