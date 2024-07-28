package stack_error

import "fmt"

type StackError interface {
	Error() string
	ErrorCode() int
	ErrorMsg() string
	String() string
}

type stackError struct {
	Code    int
	Message string
	*stack
}

func New(errCode int, message string) StackError {
	return &stackError{
		Code:    errCode,
		Message: message,
		stack:   callers(),
	}
}

func (s *stackError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, stack: %+v", s.Code, s.Message, s.stack)
}

func (s *stackError) ErrorCode() int {
	return s.Code
}

func (s *stackError) ErrorMsg() string {
	return s.Message
}

func (s *stackError) String() string {
	return s.Error()
}
