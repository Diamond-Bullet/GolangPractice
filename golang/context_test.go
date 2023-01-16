package golang

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Usage Guidance: https://blog.csdn.net/u013276277/article/details/108923912
// Introduction: https://zhuanlan.zhihu.com/p/110085652

// ContextWithTimeOut
// `WithDeadLine` has similar functionality. `TimeOut` is implemented with `DeadLine` fundamentally.
func TestContextWithTimeOut(t *testing.T) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var p int
	go func(c context.Context) {
		for {
			select {
			case <-c.Done(): // the child coroutine listens to `Done` channel.
				println("DONE")
				return
			default:
				fmt.Printf("I am working: %d\n", p)
				p++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(c)

	// time.Sleep(15*time.Second)
	time.Sleep(5 * time.Second)
	cancel() // Both timeout and calling method `cancel`, are sending a signal to the channel `Done`.
	time.Sleep(5 * time.Second)
}
