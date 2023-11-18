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
	m.ToString()
	m.User.ToString() // if we want to call its struct member's method, specify the struct member.
	println("————————————————————————")

	// instance and pointer have different method sets. however, whatever the method's receiver is, both the two can call it.
	m.ToStringPtr()
	(&m).ToString()
	println("————————————————————————")

	// method set：
	// instance: methods with receiver T
	// pointer: methods with receiver T or *T
	// when embeds S, T has all methods with receiver S
	// when embeds *S, T has all methods with receiver S or *S
	// when embeds S or *S, *T has all methods with receiver S or *S
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
//		tab  *itab // store interface type, object type and object method address.
//		data unsafe.Pointer // point to the object
//	}
//
// an interface is a set of methods or a set of types.
// interface can embed another one. they can't have method with the same name.
func TestCompareInterface(t *testing.T) {
	// type eface struct {
	//	 _type *_type
	//	 data  unsafe.Pointer
	// }
	// empty interface comprises no method, so it's implemented by any type.

	var t1, t2 interface{}
	println(t1 == t2, t1 == nil)
	// if the type set to the interface is comparable, so does the interface.
	t1, t2 = 100, 100
	println(reflect.TypeOf(t1).String())
	println(t1 == t2)

	// an interface consists of two parts: data, type. if and only if both the two parts are nil, the interface is nil。
	var a interface{} = nil
	var b interface{} = (*int)(nil) // despite no data, value set to `b` has type.
	fmt.Printf("Part 3: a == nil: %t, b == nil: %t\n", a == nil, b == nil)
}

func TestAnonymousInterface(t *testing.T) {
	// anonymous interface
	var i interface {
		ToString()
	} = User{}

	fmt.Println(i)
}

// TODO organize sections about interface
func TestInterface(t *testing.T) {
	// interface combination
	var mm StringType2
	var m StringType1 = mm
	println(m, "\n")

	// 把变量赋值给接口时，会发生复制
	// unaddressable的变量不可赋值

	// 两个方法集相同的接口，可以作比较。
	// 先比较类型，再比较方法。接口默认值是nil。


	// type conversion of interface
	// conversion between interfaces have to use `ok` pattern, otherwise panic is caused.
	if x, ok := mm.(StringType1); ok {
		println(x)
	}
	// convert interface to specified struct. you can use `ok` pattern or `switch a.(type) case int ...`, etc.
	var b interface{} = (*int)(nil)
	fmt.Println("convert interface b to *int:", b.(*int))

	// check if a struct implements an interface by compiler
	// var x string
	// var _ StringType1 = x // prompt an error, because `string` doesn't implement `StringType1`
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
	// implicit type conversion.
	// it's feasible since `TB` implements `testing.TB`
	// In this way, `testing.TB` is implemented outside its package and private methods are circumvented.
	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground")
}
