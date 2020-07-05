package math

import "testing"

func testMinus(t *testing.T) {
	value := Minus(1, 2)
	if value != -1 {
		t.Errorf("Got %d, want %d", value, -1)
	}
}

func BenchmarkMinus(b *testing.B) {
	for n := 0; b.N <= n; n++ {
		Minus(1, 2, 3, 4)
	}
}
