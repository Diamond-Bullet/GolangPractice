package stack_error

import (
	"GolangPractice/lib/logger"
	"fmt"
)

type ErrorType struct {
	Code       int
	MetricName string
}

var (
	InternalError = ErrorType{
		Code:       10000,
		MetricName: "InternalError",
	}
)

type Error interface {
	Error() string
	ErrorCode() int
	ErrorMsg() string
}

type StackError struct {
	Code    int
	Message string
	*stack
}

func New(errCode int, message string) Error {
	return &StackError{
		Code:    errCode,
		Message: message,
		stack:   callers(),
	}
}

func (s *StackError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, stack: %+v", s.Code, s.Message, s.stack)
}

func (s *StackError) ErrorCode() int {
	return s.Code
}

func (s *StackError) ErrorMsg() string {
	return s.Message
}

func NewErrWithLog(errCode int, message string) Error {
	stackError := &StackError{
		Code:    errCode,
		Message: message,
		stack:   callers(),
	}
	logger.Error(stackError.Error())
	return stackError
}

func NewWithLog(errType ErrorType, format string, v ...any) Error {
	message := fmt.Sprint(errType.MetricName+": "+format, v)

	stackError := &StackError{
		Code:    errType.Code,
		Message: message,
		stack:   callers(),
	}

	// log
	logger.Error(stackError.Error())
	return stackError
}

func NewWithLogMetric(errType ErrorType, format string, v ...any) Error {
	message := fmt.Sprint(errType.MetricName+": "+format, v)

	stackError := &StackError{
		Code:    errType.Code,
		Message: message,
		stack:   callers(),
	}

	// log
	logger.Error(message, " stack: ", stackError.Error())

	// metric
	Inc(errType)

	return stackError
}

func Inc(errType ErrorType) {
}
