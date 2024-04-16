package design_pattern

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
我们一定要牢记设计模式是前人总结的一套可以有效解决问题的经验，不要一写代码就在考虑该使用什么设计模式，这是极其不可取的。
正确的做法应该是在实现业务需求的同时，尽量遵守面向对象设计的六大设计原则即可。
后期随着业务的扩展，你的代码会不断的演化和重构，在此过程中设计模式会大放异彩的。
*/

/*
原型模式

当一个对象的构建代价过高时。例如某个对象里面的数据需要访问数据库才能拿到，而我们却要多次构建这样的对象。
当构建的多个对象，均需要处于某种原始状态时，就可以先构建一个拥有此状态的原型对象，其他对象基于原型对象来修改。
*/
type ProtoType struct{}

// 对象池模式
func TestObjectPool(t *testing.T) {
	pool := NewPool(10)
	go pool.Run()

	for i := 0; i < 40; i++ {
		pool.Wg.Add(1)
		pool.PushTask(&Task{ID: i, work: func(i ...interface{}) error {
			time.Sleep(time.Second)
			return nil
		}})
	}

	pool.Wg.Wait()
	fmt.Println("goroutines: ", runtime.NumGoroutine())
	pool.Stop()
	time.Sleep(time.Second)
	fmt.Println("goroutines: ", runtime.NumGoroutine())
}

// 构造器模式
func TestBuilder(t *testing.T) {
	mac := NewMac("Intel Core i7-10900", "Santax 16GB DDR4")
	director := NewDirector(mac)
	director.Construct("MAC Standard KeyBoard", "SkyWorth 15.6", 6)
	println(mac.GetComputer())
}
