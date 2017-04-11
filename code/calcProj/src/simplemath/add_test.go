// add_test.go
package simplemath

import "testing"

func TestAdd1(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1 ,2) failed. Got %d, expected 3.", r)
	}
}

func BenchmarkAdd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
