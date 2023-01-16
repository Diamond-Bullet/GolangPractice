package golang

import (
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"testing"

	_ "net/http/pprof"
)

// runtime.NumGoroutine() 当前协程数量

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
		log.Println(http.ListenAndServe(":61111", nil))
	}()

	// 业务代码
	select {}
}

// trace 采样一段时间，指标跟踪分析工具
func TestTrace(t *testing.T) {
	// 输出信息到pwd/trace.out
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

	// 业务代码
	select {}
}
