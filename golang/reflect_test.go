package golang

import (
	"fmt"
	"reflect"
	"testing"
)

// https://zhuanlan.zhihu.com/p/269096255 #introduce CanSet, CanAddr, flag

// Reflect.
// an Interface type variable stores a 2-elements tuple: (value, type)
func TestReflectBase(t *testing.T) {
	var x float64
	fmt.Printf("x type: %v\n", reflect.TypeOf(x))   // Type
	fmt.Printf("x value: %v\n", reflect.ValueOf(x)) // Value
	println()

	v := reflect.ValueOf(x) // type reflect.Value, some other methods can refer to
	println(v.Float())
	println()

	type X int
	var xx X = 10
	// type表示显示类型，kind表示其底层类型
	fmt.Println(reflect.TypeOf(xx).Name(), reflect.TypeOf(xx).Kind())
	println()

	// 构造基础复合类型
	// a := reflect.ArrayOf(10, reflect.TypeOf(byte(0))) [10]uint8
	// m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(int(0))) map[string]int

	// Elem() 返回指针，管道，切片等复合类型的基础类型
	p := reflect.TypeOf(&x)
	fmt.Println(p, p.Elem())
}

func TestReflectTypeOf(t *testing.T) {
	var x User
	tx := reflect.TypeOf(x)
	fmt.Println("x type: ", tx) // 类型

	fmt.Println("x package path: ", tx.PkgPath())
}

func TestReflectValueOf(t *testing.T) {
	var x int
	vx := reflect.ValueOf(x)
	fmt.Println("vx can set: ", vx.CanSet())

	// go1.16.5: Elem只能用于interface, ptr, 获取指向的基础类型
	vpx := reflect.ValueOf(&x)
	fmt.Println("vpx can set: ", vpx.CanSet()) // 是否可以赋值
	fmt.Println("vpx.Elem can set: ", vpx.Elem().CanSet())

	vpx.Elem().SetInt(20)
	println(x)

	var ix = interface{}(&x)
	vix := reflect.ValueOf(ix)
	fmt.Println("vix can set: ", vix.CanSet())
	fmt.Println("vix.Elem can set: ", vix.Elem().CanSet())
}
