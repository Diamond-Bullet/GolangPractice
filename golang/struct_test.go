package golang

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

// struct
func TestAnonymousStruct(t *testing.T) {
	// anonymous struct, `_` field
	m := struct {
		_    int
		Name string
		Age  int
	}{Name: "song", Age: 18}
	fmt.Println(m)
}

func TestCompare(t *testing.T) {
	// when all fields support operators ==、!=, this struct is comparable.
	type Foo struct {
		Name string
		Age  int
		// Child []int
	}
	f1 := Foo{Name: "song", Age: 12}
	f2 := Foo{Name: "li", Age: 12}
	f3 := Foo{Name: "li", Age: 12}
	fmt.Println("f1 == f2: ", f1 == f2)
	fmt.Println("f2 == f3: ", f2 == f3)
}

func TestEmptyStruct(t *testing.T) {
	// the length of empty struct is 0, tha same as it's an item of the array.
	// a variable does not need actual memory space in heap, will point to runtime.zerobase.
	// for instance, a slice of empty struct
	b := struct{}{}
	bs := [100]struct{}{}
	println("size of b:", unsafe.Sizeof(b))
	println("size of bs:", unsafe.Sizeof(bs))

	// use empty struct in channel to communicate between goroutines.
	c := make(chan struct{})
	go func() {
		<-time.After(3 * time.Second)
		c <- struct{}{}
	}()

	select {
	case <-c:
		println("get struct, exit")
	}
}

func TestEmbedStruct(t *testing.T) {
	type Foo struct {
		Name string
		Age  int
	}

	type Foo1 struct {
		Name string
	}

	// embedded struct

	// anonymous field, use type but not particular name. in this case, this field is named with its type name.
	// we can use members in anonymous field directly.
	// anonymous field can be any type or pointer.
	type Bar struct {
		Foo
		FooFoo  Foo1
		Height  int `height:"180"` // Tag
		int                        // name: int
		*string                    // name: string
		// since `string` and `*string` have the same field name, they can not exist together.
	}

	bb := Bar{Height: 12}
	v := reflect.ValueOf(bb)
	println(v.Type().Field(0).Tag)
}

func TestMethod(t *testing.T) {
	// an instance can call its methods and methods of its struct members.
	// if instance `m` and one of it struct member have methods with the same name `Foo`, m.Foo() is calling m's.
	m := Manager{}
	m.toString()
	m.User.toString() // if we want to call its struct member's method, specify the struct member.
	println("————————————————————————")

	// instance and pointer have different method sets. however, whatever the method's receiver is, both the two can call it.
	m.toString2()
	(&m).toString()
	println("————————————————————————")

	// method set：
	// instance: methods with receiver T
	// pointer: methods with receiver T or *T
	// when embed S, T has all methods with receiver S
	// when embed *S, T has all methods with receiver S or *S
	// when embed S or *S, *T has all methods with receiver S or *S
	ty := reflect.TypeOf(m)
	for i, n := 0, ty.NumMethod(); i < n; i++ {
		me := ty.Method(i)
		fmt.Println(me.Name, me.Type)
	}

	// var _ StringType1 = Manager{} // `Manager` does not implement the interface
	var _ StringType1 = &Manager{} // `*Manager` does
}

// interface
//
//	type iface struct {
//		tab  *itab // store interface typ, object type and object methods addresses.
//		data unsafe.Pointer // point to the object
//	}
//
// an interface is a set of methods or a set of types.
// interface can embed another one. they can't have method with the same name.
func TestInterface(t *testing.T) {
	// 空接口没有方法，所有被任何类型实现
	// 如果实现接口的类型支持，那么接口可比较
	var t1, t2 interface{}
	println(t1 == t2, t1 == nil)
	t1, t2 = 100, 100
	println(reflect.TypeOf(t1).String())
	println(t1 == t2)

	// 接口组合
	var mm StringType2
	var m StringType1 = mm
	println(m, "\n")

	// 匿名接口
	var tt interface {
		toString()
	} = User{}
	println(tt)
	println()

	// 把变量赋值给接口时，会发生复制
	// unaddressable的变量不可赋值

	// 两个方法集相同的接口，可以作比较。
	// 先比较类型，再比较方法。接口默认值是nil。

	// 接口变量的两部分都为nil, 接口才为nil。
	var a interface{} = nil
	var b interface{} = (*int)(nil) // b是有类型的
	println(a == nil, b == nil)
	println()

	// 接口的类型转换。
	// 接口和接口, 不使用ok模式会panic
	if x, ok := mm.(StringType1); ok {
		println(x)
	}
	// 接口和具体类型，同样可以ok模式，或者 switch a.(type) case int ...
	println(b.(*int))

	// 通过编译器检查是否实现某个接口
	// var x string
	// var _ StringType1 = x // 提示错误，因为x并没有实现该接口
}

// when embedding anonymous variables, TB can use all the variable's methods.
// of course, it's the anonymous variable calling the method
type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

func TestInterfaceWrapper(t *testing.T) {
	// 隐式转换，因为TB实现了testing.TB的所有方法
	// 这样就跳过了私有方法，而在外部实现了testing.TB接口。
	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground")
}
