package design_pattern

/*
单例模式
	某个类只有一个实例，且自行实例化并向整个系统提供此实例

你可以直接使用封装好的 sync.Once() 方法，实现单例模式
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
