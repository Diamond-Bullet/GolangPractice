package error

import (
	"GolangPractice/pkg/logger"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(-10000, "good")
	logger.Error(err)
}
