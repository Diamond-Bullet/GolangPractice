package golang

import (
	"GolangPractice/utils/logger"
	"net/http"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"testing"

	_ "net/http/pprof"
)

/*
	go tool 查看内置工具

	nm：查看 二进制文件 的符号表（等同于系统 nm 命令）

	objdump：反汇编工具，分析二进制文件（等同于系统 objdump 命令）

	cover：生成代码覆盖率

	compile：代码汇编
		go tool compile -N -l -S example.go 禁用内联和代码优化
		go tool compile -S example.go   查看汇编输出
*/

// runtime.NumGoroutine() 当前协程数量

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
		logger.Infoln(http.ListenAndServe(":61111", nil))
	}()

	// your product code here
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

	// your product code here
	select {}
}
