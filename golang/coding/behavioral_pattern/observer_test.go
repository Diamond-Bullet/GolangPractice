package behavioral_pattern

import (
	"github.com/gookit/color"
	"reflect"
	"testing"
	"time"
)

/*
Observer, also named Publish-Subscribe.
It defines a 1-to-n dependency among objects. When the state of an object changes, all subscribers get notified and updated.
Push, Pull.

Applicable scenarios:
1. There are several classes cooperating with each other in a system, which we need to maintain the consistency of them without introducing too much coupling.
2. The change of one object leads to changes of multiple other objects.
*/

func TestObserver(t *testing.T) {
	clock := &Clock{
		Observers: make(map[Observer]struct{}),
		Time:      time.Now().String(),
	}

	digitalGUI := &DigitalGUI{}
	clock.Attach(digitalGUI)

	clock.Update()
}

type Subject interface {
	Attach(o Observer)
	Detach(o Observer)
	Update()
	notify()
	GetMsg() string
}

type Observer interface {
	Update(s Subject)
}

type Clock struct {
	Observers map[Observer]struct{}
	Time      string
}

func (c *Clock) Attach(o Observer) {
	c.Observers[o] = struct{}{}
}

func (c *Clock) Detach(o Observer) {
	delete(c.Observers, o)
}

func (c *Clock) Update() {
	c.Time = time.Now().String()
	c.notify()
}

func (c *Clock) GetMsg() string {
	return c.Time
}

func (c *Clock) notify() {
	for o := range c.Observers {
		o.Update(c)
	}
}

type DigitalGUI struct {
	Time string
}

func (d *DigitalGUI) Update(s Subject) {
	d.Time = s.GetMsg()
	color.Blueln(d.Time)
}

/*
Now we add a layer ChangeManager between Subject and Observers.
1. 维护一个目标的观察者集合
2. 定义一个特定的更新策略
3. 根据一个目标的请求，更新所有依赖于这个目标的观察者
*/

type ComplexClock struct {
	Time string
	Date string
	ChangeManager
}

func (c *ComplexClock) Attach(o Observer) {
	c.ChangeManager.Register(c, o)
}

func (c *ComplexClock) Detach(o Observer) {
	c.ChangeManager.Unregister(c, o)
}

func (c *ComplexClock) Update() {
	c.Time = time.Now().String()
	c.notify()
}

func (c *ComplexClock) notify() {
	c.ChangeManager.Notify(c)
}

func (c *ComplexClock) GetMsg() string {
	return c.Time
}

type ChangeManager interface {
	Register(s Subject, o Observer)
	Unregister(s Subject, o Observer)
	Notify(s Subject)
}

type DefaultChangeManager struct {
	RegisterInfo map[Subject][]Observer
}

func (d *DefaultChangeManager) Register(s Subject, o Observer) {
	d.RegisterInfo[s] = append(d.RegisterInfo[s], o)
}

func (d *DefaultChangeManager) Unregister(s Subject, o Observer) {
	registerList, ok := d.RegisterInfo[s]
	if ok {
		for i, oo := range registerList {
			if reflect.DeepEqual(oo, o) {
				registerList = append(registerList[:i], registerList[i+1:]...)
			}
		}
		d.RegisterInfo[s] = registerList
	}
}

// Notify 定制通知策略，随机通知一个
func (d *DefaultChangeManager) Notify(s Subject) {
	for _, o := range d.RegisterInfo[s] {
		o.Update(s)
		return
	}
}
