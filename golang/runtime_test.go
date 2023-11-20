package golang

import (
	"runtime"
	"runtime/debug"
	"testing"
)

// runtime stands for the time the go process is running.

// runtime.NumCPU() the physical machine's amount of CPU.
// runtime.GOMAXPROCS() Set Processor `P`'s count.
// runtime.Gosched() Let current goroutine give up execution right, so other goroutines can be scheduled.
// runtime.Goexit() Terminate present goroutine. all `defer` registered before are supposed to be executed. if it's called in main.main, the program will panic after all tasks is finished.

func TestMajorFunc(t *testing.T) {
	runtime.GC() // garbage collection

	runtime.NumGoroutine() // current goroutine amount
}

func TestDebug(t *testing.T) {
	debug.PrintStack() // print calling stack
}
