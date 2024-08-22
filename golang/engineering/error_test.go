package engineering

import (
	"GolangPractice/utils/logger"
	"errors"
	"fmt"
	pkgerrors "github.com/pkg/errors"
	"testing"
)

// provided by Go team. simply wrap an error with new prefix.
func TestWrapError(t *testing.T) {
	err := errors.New("error here")
	err1 := fmt.Errorf("layer1: %w", err)
	err2 := fmt.Errorf("layer2: %w", err1)
	logger.Errorln(err2)
}

func StackError1() error {
	return StackError2()
}

func StackError2() error {
	return pkgerrors.New("error here")
}

// https://github.com/pkg/errors error with stack trace
// alternative: https://github.com/go-errors/errors
// learn about new error handling draft `Go2 errors` by Go team.
func TestStackError(t *testing.T) {
	err := pkgerrors.Errorf("err: %s", "i want to bring out an error")
	logger.Infoln(err)
	logger.Infof("%+v", err)

	err1 := pkgerrors.Wrap(err, "err1")
	logger.Infof("%+v", err1)

	err2 := StackError1()
	logger.Infof("%+v", err2)
}

type RpcError[T any] struct {
	Code  int64
	Param T
}

func (m *RpcError[T]) Error() string {
	return fmt.Sprintf("rpc error code: %d, msg: %v", m.Code, m.Param)
}
