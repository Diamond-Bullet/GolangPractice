package golang

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

//some famous open-source libraries
// https://github.com/bouk/monkey author bouke from Netherlands
// https://github.com/agiledragon/gomonkey author Xiaolong Zhang from China

// It seems like the difference between the two repo is tiny. Or the latter might have some improvement from the former.
// you might add `-gcflags "all=-N -l"` when run `go test`, or the patch can not work correctly.
// explain: -N disable optimizations, -l disable inline.
// run `go tool compile -help` to get more about `gcflag`.

// go test -v  golang.go -test.run Add #Test the single func `TestAdd`
// go test -v  golang.go -test.bench ForFun -test.run ForFun #test a func in benchmark mode
// go test -bench=. -benchtime=3s #run all benchmarks, specify the test duration.
// -cpuprofile profile.out #output the result of cpu analysis
// -memprofile memprofile.out #Output the result of Mem analysis
// https://blog.csdn.net/weixin_34232617/article/details/91854391

/* An example for another way to use `go test`. And it's related to memory allocation from go101.
var t *[5]int64
var s []byte

func f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = &[5]int64{}
	}
}

func g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = make([]byte, 32769)
	}
}

func main() {
	println(unsafe.Sizeof(*t))      // 40
	rf := testing.Benchmark(f)
	println(rf.AllocedBytesPerOp()) // 48
	rg := testing.Benchmark(g)
	println(rg.AllocedBytesPerOp()) // 40960
}
*/

// stress test
func BenchmarkDirect(b *testing.B) {
	x, y := 1, 2
	for i := 0; i < b.N; i++ {
		_ = x + y
	}
}

// bulk-cases test
func TestAdd(t *testing.T) {
	var tests = []struct {
		Name   string
		ArgA   int
		ArgB   int
		Result int
		Err    bool
	}{
		{"case1", 1, 2, 3, false},
		{"case2", 3, -2, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := Add(tt.ArgA, tt.ArgB)
			if got != tt.Result {
				t.Errorf("Add() = %v, want = %v", got, tt.Result)
			}
		})
	}
}

// most kinds of mock
func TestMock(t *testing.T) {
	// Mock Function
	// make it by replace the func address
	convey.Convey("test", t, func() {
		patches := gomonkey.ApplyFunc(Add, func(a, b int) int {
			return 1
		})
		defer patches.Reset()

		convey.So(Add(1, 3), convey.ShouldEqual, 1) // pass
		convey.So(Add(1, 3), convey.ShouldEqual, 4) // fail
	})
}

func Add(x, y int) int {
	return x + y
}
