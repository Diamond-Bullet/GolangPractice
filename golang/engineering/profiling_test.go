package engineering

import (
	"GolangPractice/pkg/logger"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"runtime/trace"
	"testing"

	_ "net/http/pprof"
)

// `go tool`: list embedded tools
//
// `go tool nm main.o`：view symbols in an object file (similar to `nm` in os).
// to generate an object file, use `go build -o main.o main.go`.
//
// `go tool objdump -S main.o`：disassemble an object file to source code. like `objdump` in os.
//
// `go tool cover -html=coverage.out -o coverage.html`: generate an HTML coverage report
//
// `go tool compile -N -l -S example.go`: disable inlining and other code optimization.
// `compile` generate an `.o` file. it is typically invoked by `go build`.

func TestRuntimeFuncs(t *testing.T) {
	runtime.GC() // garbage collection

	debug.PrintStack() // print call stack

	runtime.NumCPU() // the physical machine's amount of CPU.

	runtime.NumGoroutine() // current goroutine amount

	runtime.GOMAXPROCS(10) // Set Processor `P`'s count.

	runtime.Gosched() // Let current goroutine give up execution right, so other goroutines can be scheduled.

	// Terminate present goroutine. all `defer` registered before are supposed to be executed.
	// if it's called in main.main, the program will panic after all tasks is finished.
	runtime.Goexit()
}

// `go tool pprof [pprof_file]` to enter interactive mode.
// use help to get commands you can use, like `top 5`, `list [func_name]`, `png`, etc.
// to visualize profile, install `graphviz` using command like `yum install graphviz`.
// for more at https://go.dev/blog/pprof
func TestRuntimePprof(t *testing.T) {
	// cpu
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// 堆内存 pprof.WriteHeapProfile()
}

// net/http/pprof
func TestNetHttpPprof(t *testing.T) {
	// pprof
	go func() {
		logger.Info(http.ListenAndServe(":61111", nil))
	}()

	// your product code here
	select {}
}

// `runtime/trace` samples for a period of time. trace and analyze the execution of program.
// especially useful for diagnosing performance issues.
func TestTrace(t *testing.T) {
	// output information to $(pwd)/trace.out
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// your business code here
}
