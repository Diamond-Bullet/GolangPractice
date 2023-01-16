package golang

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"
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
func Test1(t *testing.T) {
	innerFunc := Foo()
	//println(&a)
	//fmt.Printf("a size: %d\n", unsafe.Sizeof(a))
	//println(&b)
	//fmt.Printf("b size: %d\n", unsafe.Sizeof(b))
	//println(&innerFunc)
	//fmt.Printf("innerFunc size: %d\n", unsafe.Sizeof(innerFunc))
	//println(pA)
	//fmt.Printf("pA size: %d\n", unsafe.Sizeof(pA))
	println(innerFunc(1))
	println(innerFunc(2))
}

func Foo() func(x int) int {
	a := 1
	return func(x int) int {
		a += x
		return a
	}
}

// memory alignment
// https://blog.csdn.net/u011957758/article/details/85059117
func TestMemoryAlignment(t *testing.T) {
	type Part2 struct {
		Name int32
		Age  *int32
	}

	type Part1 struct {
		a bool
		b int32
		c int8
		d int64
		Part2
		e byte
	}

	p := Part1{}
	var x = 1<<63 - 1
	println(unsafe.Offsetof(p.e))
	fmt.Printf("p align: %d\n", unsafe.Alignof(p.e))
	println(&(p.e))
	println(&x)
	fmt.Printf("p size: %d\n", unsafe.Sizeof(p))
}

// pointer computation/operation/calculation
func TestPointer(t *testing.T) {
	type S1 struct {
		A int32
		B int64
	}

	s := S1{}
	fmt.Println(s)
	b := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.B)))
	*b = 1
	fmt.Println(s)
}

// defer: last in, first out
// reference：https://studygolang.com/articles/16067
//		https://www.cnblogs.com/makelu/p/11226974.html
func TestDefer(t *testing.T) {
	defer func() {
		println(1)
	}()

	defer func() {
		println(2)
	}()
	// output
	// 2
	// 1
}

func TestDefer1(t *testing.T) {
	var x int
	defer func() {
		fmt.Printf("x == 0: %v\n", x == 0) // false # as wrapped with an anonymous func, the variable is not certain until defer func really executes.
	}()
	defer fmt.Printf("x == 1: %v\n", x == 1) //false # the value is certain right now.

	x++
}

// keyword 'return' actually has two steps in golang.
func TestDefer2(t *testing.T) {
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
	println(tt) // 1
}

// recover panic
// `recover` can only heal panic in current goroutine, since every goroutine has its own stack.
func TestRecover(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from crash: %v", r)
		}
	}()

	panic("I want to panic")
}

// no captured/caught panic will crash the main progress.
func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from crash: %v", r)
		}
	}()
	//go func() {
	//	JustForPanic()
	//}()

	JustForPanic()

	time.Sleep(time.Second)
}

func JustForPanic() {
	panic("I want to PANIC")
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
			println(*px)
		}
	}
}

// string runtime/string.go
//	type stringStruct struct {
//		str unsafe.Pointer
//		len int
//	}
func TestString(t *testing.T) {
	s := "abcdefg 好东西 @#¥" // pointer points to a bytes array.
	s1 := s[2:4]

	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
	println()

	// byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %c  ", i, s[i])
	}
	println()

	// rune
	// library `utf8` offers some functions handling type `rune`.
	for i, x := range s {
		fmt.Printf("%d: %c  ", i, x)
	}
	println("\n")

	// In general(normally, usually, commonly, generally), modifying string has to transform it into []byte、[]rune, which calls a reallocation of memory and data duplication.
	// If we just need to turn string into []byte in a read-only situation, using unsafe.Pointer instead can avoid extra actions mentioned above.
	// TODO how to alter string by unsafe.Pointer
	b := []byte("1234")
	sb := *(*string)(unsafe.Pointer(&b))
	bs := *(*[]byte)(unsafe.Pointer(&sb))
	bs[1] = 'a'
	println(sb)
	fmt.Printf("%s\n", bs)

	// Concatenating/Joining/Splicing strings dynamically needs to reallocate memory constantly, leading to(resulting in) worse performance.
	// Allocating memory in advance(ahead, beforehand).
	// strings.Join() will statistics required room, and apply at one time.
	// bytes.NewBufferString("") is most efficient.
}

// array
func TestArray(t *testing.T) {
	d1 := [4]int{3: 10}   // Initialize by the index.
	d2 := [...]int{2: 10} // define/clear/crystallize `length` characteristic/property according to the biggest index has been declared by the user.
	fmt.Printf("%v\n", d1)
	fmt.Printf("%v\n", d2)

	d3 := [3][3]int{}     // multidimensional array.
	d4 := [...][3]int{{}} // the characteristic `[...]` is available if and only if it's the first layer.
	fmt.Printf("%v\n", d3)
	fmt.Printf("%v\n", d4)

	// if the array's basic(underlying) type supports operations like `==`、`!=`, the array does too(then so does the array).
	a, b := [2]int{1, 2}, [2]int{2, 4}
	println(a == b)
}

// slice
func TestSlice(t *testing.T) {
	var s1 []int
	s2 := []int{}
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s1))) // 0x0, invalid.
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s2))) // s2 is initialized, and the pointer is set, but point to `zerobase`
	println()
	// by slice, it's easy to implement stack and queue.

	// `append`, will add new element to the end of the underlying array.
	s := make([]int, 0, 10)
	s3 := append(s, 10)
	s4 := append(s, 20)
	fmt.Println(s3)
	fmt.Println(s4)
	// when capacity is no longer sufficient, slice will scale up，a new underlying array will be assigned to replace the old and smaller one.
	s5 := s[:2:3] // start:end:cap
	s6 := append(s5, 20, 30, 40, 50)
	fmt.Println(&s5[0])
	fmt.Println(&s6[0])
}

// map
// https://blog.csdn.net/u010853261/article/details/99699350
func TestMap(t *testing.T) {
	m := map[int]int{1: 1}

	change := func(mm map[int]int) {
		mm[1] = 2
	}

	change(m)
	println(m[1]) // result:2, map is a pointer

	// adding or eliminating key during traversing the map is safe.
	for i := 2; i < 5; i++ {
		m[i] = i + 10
	}
	for k := range m {
		m[k+10] = k + 20
		delete(m, k+1)
		fmt.Println(k, m)
	}

	// Can not do some other concurrent operations at the same time when writing the map, unless using the lock or concurrency-safe data structures, like sync.Map。
	// `fatal: concurrent write\read` will occur when you really start the program.

	// prepare appropriate room in advance，to avoid reallocating memory and rehashing after scaling up the mao.
}

// Currently, after deleting enough elements from map，it won't shrink automatically, referring to `src/runtime/map.go`
func TestMapShrinkWhenDelete(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 10000; i++ {
		m[i] = i
		if (i+1)%100 == 0 {
			// this expression is based on the structure of map, map returned by `make` is a reference to `hmap`
			// you can see the comprised fields of it in the mentioned file.
			fmt.Print(*(*uint8)(unsafe.Pointer(uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&m))) + uintptr(9))), " ")
		}
	}
	println()
	for i := 0; i < 10000; i++ {
		delete(m, i)
		if (i+1)%1000 == 0 {
			fmt.Print(*(*uint8)(unsafe.Pointer(uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&m))) + uintptr(9))), " ")
		}
	}
}
