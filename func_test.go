package utils

import (
	"errors"
	"testing"
)

func TestSafeRun(t *testing.T) {
	expectErr := errors.New("should error")
	fn := func() { panic(expectErr) }
	err := SafeRun(fn)
	if err != expectErr {
		t.Error("SafeRun")
	}
}
