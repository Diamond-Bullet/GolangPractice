package engineering

import (
	"GolangPractice/utils/logger"
	"github.com/gookit/color"
	"reflect"
	"runtime"
	"runtime/debug"
	"testing"
	"time"
	"unicode/utf8"
	"unsafe"
)

// 《Go语言学习笔记》

/*
$GOPATH /pkg /src /bin

GOPATH: where to store the third party libraries
go get: download libraries

GOROOT: go source code and tool chain

GOBIN: executable files
*/

/*
Global variables initialization
init(): all functions named `init()` execute in a wrapped function in runtime. every file can define multiple `init()`.
main(): your program's entrance
*/

// Internal package `internal`
// every package belongs to the `internal` can visit each other.
// package `internal` can be visited by the packages belong to the same parent package, but is invisible to external packages.

// go version

// the way `Closure` references a variable outside the anonymous function is pointer.
func TestClosure(t *testing.T) {
	Foo := func() func(x int) int {
		a := 1
		return func(x int) int {
			a += x
			return a
		}
	}

	innerFunc := Foo()

	logger.Infoln(&innerFunc)
	logger.Infof("innerFunc size: %d", unsafe.Sizeof(innerFunc))
	logger.Infoln(innerFunc(1)) // 2
	logger.Infoln(innerFunc(2)) // 4
}

// memory alignment
// https://blog.csdn.net/u011957758/article/details/85059117
func TestMemoryAlignment(t *testing.T) {
	type Part1 struct {
		Name int32
		Age  *int32
	}

	type Part2 struct {
		a bool
		b int32
		c int8
		d int64
		Part1
		e byte
	}

	p := Part2{}
	var x = 1<<63 - 1
	logger.Infoln(unsafe.Offsetof(p.e))
	logger.Infof("p align: %d", unsafe.Alignof(p.e))
	logger.Infoln(&(p.e))
	logger.Infoln(&x)
	logger.Infof("p size: %d", unsafe.Sizeof(p))
}

// pointer computation/operation/calculation
func TestPointer(t *testing.T) {
	type S1 struct {
		A int32
		B int64
	}

	s := S1{}
	logger.Infoln(s)
	b := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.B)))
	*b = 1
	logger.Infoln(s)
}

// defer: last in, first out
//
// reference：https://studygolang.com/articles/16067
//
//	https://www.cnblogs.com/makelu/p/11226974.html
func TestDeferSequence(t *testing.T) {
	// output
	// 2
	// 1

	defer func() {
		logger.Infoln(1)
	}()

	defer func() {
		logger.Infoln(2)
	}()
}

func TestDeferParameter(t *testing.T) {
	var x int
	defer func() {
		logger.Infof("x == 0: %v", x == 0) // false # as wrapped with an anonymous func, the variable is not certain until defer func really executes.
	}()
	defer logger.Infof("x == 1: %v", x == 1) //false # the value is certain right now.

	x++
}

// keyword 'return' actually has two steps in golang.
func TestDeferReturn(t *testing.T) {
	tf := func() (x int) {
		var i int
		defer func() {
			x++ // take effect，modify the returned value directly
		}()
		defer func() {
			i++ // just change i, but ret = i has been executed before
		}()
		return i // x = i, defer func, ret x
	}
	tt := tf()
	logger.Infoln(tt) // 1
}

func JustForPanic() {
	panic("I want to PANIC")
}

// `recover` can only recover panic in current goroutine, since every goroutine has its own stack.
// non-captured/uncaught panic will make the main progress crash.
func TestPanicNotCaptured(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			logger.Infof("recovered from crash: %v", r)
		}
	}()

	go func() {
		JustForPanic()
	}()

	time.Sleep(time.Second)
}

func TestPanicCaptured(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			logger.Infof("recovered from crash: %v", r)                                    // here is only a line of error
			logger.Infof("recovered from crash: %v, detailed stack: %s", r, debug.Stack()) // print goroutine stack
		}
	}()

	JustForPanic()

	time.Sleep(time.Second)
}

// TODO For `Go`, memory address can change in runtime
func TestMemoryAddress(t *testing.T) {
	x := make([]int, 1, 1)
	x[0] = 1
	p := uintptr(unsafe.Pointer(&x))

	for i := 0; i < 100; i++ {
		runtime.GC()

		px := (*[]int)(unsafe.Pointer(p))
		if (*px)[0] != x[0] {
			logger.Infoln(*px)
		}
	}
}

// string runtime/string.go
//
//	type stringStruct struct {
//		str unsafe.Pointer
//		len int
//	}
func TestString(t *testing.T) {
	s := "abcdefg 好东西 @#¥" // pointer points to a bytes array.
	s1 := s[2:4]

	logger.Infof("%#v", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	logger.Infof("%#v", (*reflect.StringHeader)(unsafe.Pointer(&s1)))

	// byte
	for i := 0; i < len(s); i++ {
		logger.Infof("%d: %c  ", i, s[i])
	}

	// rune
	// library `utf8` offers some functions handling type `rune`.
	for i, x := range s {
		logger.Infof("%d: %c  ", i, x)
	}

	// In general, modifying string has to transform it into []byte、[]rune, which calls a reallocation of memory and data duplication.
	// If we just need to turn string into []byte in a read-only situation, using unsafe.Pointer instead can avoid extra actions mentioned above.
	// TODO how to alter string by unsafe.Pointer
	b := []byte("1234")
	sb := *(*string)(unsafe.Pointer(&b))
	bs := *(*[]byte)(unsafe.Pointer(&sb))
	bs[1] = 'a'
	logger.Infoln(sb)
	logger.Infof("%s", bs)

	// Concatenating/Joining/Splicing strings dynamically needs to reallocate memory constantly, leading to(resulting in) worse performance.
	// Allocating memory in advance(ahead, beforehand).
	// strings.Join() will statistics required room, and apply at one time.
	// bytes.NewBufferString("") is most efficient.
}

func TestStrLen(t *testing.T) {
	s := "abcdefg"
	logger.Infoln("len:", len(s)) // 7

	s1 := "abc好东西"
	logger.Infoln("len:", len(s1))                                          // 12
	logger.Infoln("rune len:", utf8.RuneCountInString(s1), len([]rune(s1))) // 6
}

// array
func TestArray(t *testing.T) {
	d1 := [4]int{3: 10}   // Initialize by the index.
	d2 := [...]int{2: 10} // define/clear/crystallize `length` characteristic/property according to the biggest index has been declared by the user.
	logger.Infof("%v", d1)
	logger.Infof("%v", d2)

	d3 := [3][3]int{}     // multidimensional array.
	d4 := [...][3]int{{}} // the characteristic `[...]` is available if and only if it's the first layer.
	logger.Infof("%v", d3)
	logger.Infof("%v", d4)

	// if the array's basic(underlying) type supports operations like `==`、`!=`, the array does too(then so does the array).
	a, b := [2]int{1, 2}, [2]int{2, 4}
	logger.Infoln(a == b)
}

// slice
// by slice, it's easy to implement stack and queue.
func TestSlice(t *testing.T) {
	var s1 []int
	logger.Infof("%#v", (*reflect.SliceHeader)(unsafe.Pointer(&s1))) // 0x0, invalid.

	s2 := []int{}
	logger.Infof("%#v", (*reflect.SliceHeader)(unsafe.Pointer(&s2))) // s2 is initialized, and the pointer is set, but point to `zerobase`

	// `append`, will add new element to the end of the underlying array.
	s := make([]int, 0, 10)
	s3 := append(s, 10)
	s4 := append(s, 20)
	logger.Infoln(s3)
	logger.Infoln(s4)

	// when capacity is no longer sufficient, slice will scale up，a new underlying array will be assigned to replace the old and smaller one.
	s5 := s[:2:3] // start:end:cap
	s6 := append(s5, 20, 30, 40, 50)
	logger.Infoln(&s5[0])
	logger.Infoln(&s6[0])
}

// map
// https://blog.csdn.net/u010853261/article/details/99699350
func TestMap(t *testing.T) {
	m := map[int]int{1: 1}

	change := func(mm map[int]int) {
		mm[1] = 2
	}

	change(m)
	logger.Infoln(m[1]) // result: 2, map is a pointer

	for i := 2; i < 5; i++ {
		m[i] = i + 10
	}
	// adding or eliminating key during traversing the map is safe.
	for k := range m {
		m[k+10] = k + 20
		delete(m, k+1)
		logger.Infoln(k, m)
	}

	// Can not do some other concurrent operations at the same time when writing the map,
	// unless using the lock or concurrency-safe data structures, like sync.Map。
	// `fatal: concurrent write\read` will occur when you run the program.

	// prepare appropriate room in advance，to avoid reallocating memory and rehashing after scaling up the map.
}

// Currently, after deleting enough elements from map，it won't shrink automatically, referring to `src/runtime/map.go`
func TestMapShrinkWhenDelete(t *testing.T) {
	const EntryAmount = 10000

	m := make(map[int]int)
	for i := 0; i < EntryAmount; i++ {
		m[i] = i
		if (i+1)%100 == 0 {
			// this expression is based on the structure of map, map returned by `make` is a reference to `hmap`
			// you can see the comprised fields of it in the mentioned file.
			color.Redp(*(*uint8)(unsafe.Pointer(uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&m))) + uintptr(9))), " ")
		}
	}

	for i := 0; i < EntryAmount; i++ {
		delete(m, i)
		if (i+1)%1000 == 0 {
			color.Greenp(*(*uint8)(unsafe.Pointer(uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&m))) + uintptr(9))), " ")
		}
	}
}
