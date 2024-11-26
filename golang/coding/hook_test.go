package coding

import (
	"GolangPractice/lib/logger"
	"reflect"
)

/*
	Hook is a way to add functions at specific points of a procedure.

	Here we use GORM, a typical example to demonstrate how to use hooks.
*/

// Hooks constants
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

type Session struct {
	Table interface{}
}

// CallMethod calls the registered hooks
func (s *Session) CallMethod(method string, value interface{}) {
	fm := reflect.ValueOf(s.Table).MethodByName(method)
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				logger.Error(err)
			}
		}
	}
	return
}

func (s *Session) Query() error {
	s.CallMethod(BeforeQuery, nil)

	// ...
	queryResult := []string{"result1", "result2"}

	for _, result := range queryResult {
		dest := reflect.ValueOf(result)
		s.CallMethod(AfterQuery, dest.Addr().Interface())
	}

	return nil
}
