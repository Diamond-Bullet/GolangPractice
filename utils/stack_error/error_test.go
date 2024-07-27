package stack_error

import (
	"GolangPractice/utils/logger"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(-10000, "good")
	logger.Errorln(err)
}
