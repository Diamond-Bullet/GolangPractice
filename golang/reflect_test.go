package golang

import (
	"GolangPractice/utils/logger"
	"reflect"
	"testing"
)

// https://zhuanlan.zhihu.com/p/269096255 introduce CanSet, CanAddr, flag

// Reflect.
// an Interface type variable stores a 2-elements tuple: (value, type)

func TestBase(t *testing.T) {
	var x float64
	logger.Infof("x type: %v\n", reflect.TypeOf(x))   // Type
	logger.Infof("x value: %v\n", reflect.ValueOf(x)) // Value

	v := reflect.ValueOf(x) // type reflect.Value, some other methods can refer to
	logger.Infoln(v.Float())

	type X int
	var xx X = 10
	// type表示显示类型，kind表示其底层类型
	logger.Infoln(reflect.TypeOf(xx).Name(), reflect.TypeOf(xx).Kind())

	// 构造基础复合类型
	// a := reflect.ArrayOf(10, reflect.TypeOf(byte(0))) [10]uint8
	// m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(int(0))) map[string]int

	// Elem() 返回指针，管道，切片等复合类型的基础类型
	p := reflect.TypeOf(&x)
	logger.Infoln(p, p.Elem())
}

func TestReflectType(t *testing.T) {
	var x User
	tx := reflect.TypeOf(x)
	logger.Infoln("x type: ", tx) // the type of `x`

	logger.Infoln("x package path: ", tx.PkgPath())
}

func TestReflectMethod(t *testing.T) {
	var x User
	vx := reflect.ValueOf(x)

	// argument for `Call` is the input of this method.
	// return value of `Call` is the result of this method.
	res := vx.Method(2).Call([]reflect.Value{})
	logger.Infoln(res)
	vx.MethodByName("ToStringPtr").Call([]reflect.Value{})
}

func TestReflectValue(t *testing.T) {
	var x int
	vx := reflect.ValueOf(x)
	logger.Infoln("vx can set: ", vx.CanSet())

	// go1.16.5: Elem只能用于interface, ptr, 获取指向的基础类型
	vpx := reflect.ValueOf(&x)
	logger.Infoln("vpx can set: ", vpx.CanSet()) // 是否可以赋值
	logger.Infoln("vpx.Elem can set: ", vpx.Elem().CanSet())

	vpx.Elem().SetInt(20)
	logger.Infoln(x)

	var ix = interface{}(&x)
	vix := reflect.ValueOf(ix)
	logger.Infoln("vix can set: ", vix.CanSet())
	logger.Infoln("vix.Elem can set: ", vix.Elem().CanSet())
}
