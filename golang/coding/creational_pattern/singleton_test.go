package creational_pattern

/*
Singleton
*/

import (
	"github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	convey.Convey("TestSingleton", t, func() {
		instance1 := GetInstance()
		instance2 := GetInstance()
		convey.So(instance1, convey.ShouldEqual, instance2)
	})
}

type Singleton struct{}

var (
	once     sync.Once
	instance *Singleton
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
	})
	return instance
}
