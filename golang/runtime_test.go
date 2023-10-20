package golang

import (
	"runtime"
	"runtime/debug"
	"testing"
)

// runtime stands for the time the go process is running.

func TestMajorFunc(t *testing.T) {
	runtime.GC() // garbage collection

	runtime.NumGoroutine() // current goroutine amount
}

func TestDebug(t *testing.T) {
	debug.PrintStack() // print calling stack
}
