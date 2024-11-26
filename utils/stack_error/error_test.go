package stack_error

import (
	"GolangPractice/lib/logger"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(-10000, "good")
	logger.Error(err)
}
