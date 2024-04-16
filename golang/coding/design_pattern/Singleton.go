package design_pattern

/*
单例模式
	某个类只有一个实例，且自行实例化并向整个系统提供此实例

你可以直接使用封装好的 sync.Once() 方法，实现单例模式
*/

import (
	"sync"
	"sync/atomic"
)

// Once 线程安全
type Once struct {
	Done uint32
	sync.Mutex
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.Done) == 1 {
		return
	}

	o.Lock()
	defer o.Unlock()

	if o.Done == 0 {
		defer atomic.AddUint32(&o.Done, 1)
		f()
	}
}

type SingleExample struct{}

var (
	once     Once
	instance *SingleExample
)

func Instance() *SingleExample {
	once.Do(func() {
		instance = new(SingleExample)
	})
	return instance
}
