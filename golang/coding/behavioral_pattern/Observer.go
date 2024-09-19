package behavioral_pattern

import (
	"reflect"
	"time"
)

/*
观察者模式，  发布-订阅(Publish - Subscribe)
定义对象间的一种一对多的依赖关系 ,当一个对象的状态发生改变时 , 所有依赖于它的对象都得到通知并被自动更新。
推，拉。

适用场景：
一个系统中有多个相互协作的类，需要维护其一致性，我们并不总希望为了维护一致性而导致紧密的耦合。
一个抽象模型的一个方面，被另外若干个方面依赖。一个对象的改变，将导致其他若干个对象的改变。

角色：
目标，提供注册和删除观察者的接口；	观察者，收到通知并更新的接口
*/

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

// Update 目标状态更新，随后通知观察者
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
	println(d.Time)
}

/*
现在考虑在观察者和目标之间增加一个 变更管理器 ChangeManager
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
