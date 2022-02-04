package go_utils

import "testing"

func TestLogger(t *testing.T) {
	if got := InitLogger(".", "test"); got == nil {
		t.Errorf("InitLogger('.', test) returned nil")
	}
}
