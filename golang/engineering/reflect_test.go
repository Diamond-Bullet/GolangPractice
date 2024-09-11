package engineering

import (
	"GolangPractice/utils/logger"
	"reflect"
	"testing"
)

// https://zhuanlan.zhihu.com/p/269096255 introduce CanSet, CanAddr, flag

// reflect.Type provide methods to query type information of a struct.
// like type, package path, fields' types and names.
func TestReflectType(t *testing.T) {
	var x User

	reflectTypeUser := reflect.TypeOf(x)

	logger.Infoln("x type: ", reflectTypeUser)                   // golang.User
	logger.Infoln("x package path: ", reflectTypeUser.PkgPath()) // GolangPractice/golang

	// reflect.Type is more comprehensive, providing specific type information.
	// reflect.Kind is kind of a rough classification, telling underlying type like whether it is an int, map, struct, etc.
	logger.Infoln(reflectTypeUser.Name(), reflectTypeUser.Kind())
}

func TestReflectValue(t *testing.T) {
	var x User
	reflectValueUser := reflect.ValueOf(x)

	if reflectValueUser.CanSet() {
		reflectValueUser.Set(reflect.ValueOf(User{}))
	}

	// Elem is applicable to interfaces, pointers. To get the value the interface contains or the pointer points to.
	if reflectValueUser.Elem().CanSet() {
		reflectValueUser.Elem().Set(reflect.ValueOf(User{}))
	}
}

func TestReflectValueMethod(t *testing.T) {
	var x User
	reflectValueUser := reflect.ValueOf(x)

	if reflectValueUser.Type().Name() != "golang.User" {
		return
	}

	// argument for `Call` is the input of this method.
	// return value of `Call` is the result of this method.
	res := reflectValueUser.Method(2).Call([]reflect.Value{})
	logger.Infoln(res)

	reflectValueUser.MethodByName("ToStrParam").Call([]reflect.Value{reflect.ValueOf("Fabulous!")})
}

func TestReflectValueField(t *testing.T) {
	x := User{Name: "Alice", Age: "18"}
	reflectValueUser := reflect.ValueOf(x)

	field := reflectValueUser.FieldByName("Name")
	if field.CanSet() {
		field.SetString("123")
	}
	logger.Infoln(field.String())

	numField := reflectValueUser.Field(2)
	if numField.CanSet() {
		numField.SetString("456")
	}
}

// TODO mechanisms and applications
func TestCompoundType(t *testing.T) {
	arr := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))                  // [10]uint8
	mapping := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(int(0))) // map[string]int

	logger.Infoln(arr)
	logger.Infoln(mapping)
}
