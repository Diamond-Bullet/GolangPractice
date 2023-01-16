package golang

import (
	"runtime"
	"runtime/debug"
	"testing"
)

// runtime go运行时

func TestMajorFunc(t *testing.T) {
	runtime.GC() // 垃圾回收

	runtime.NumGoroutine() // 当前goroutine数量
}

func TestDebug(t *testing.T) {
	debug.PrintStack() // 打印调用栈
}
