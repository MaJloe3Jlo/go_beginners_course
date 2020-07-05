package math

import "testing"

func TestPlus(t *testing.T) {
	value := Plus(1, 2)
	if value != 3 {
		t.Errorf("Got %d, want %d", value, 3)
	}
}
