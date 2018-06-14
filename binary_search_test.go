package search

import (
	"log"
	"testing"
)

var ln = log.Println

func TestBinarySearch(t *testing.T) {
	s := make([]int, 100)
	for i := 0; i < 100; i++ {
		s[i] = i
	}
	a := BinarySearch(s, 2)
	ln(a)
	if a < 0 {
		t.Error("not found")
	}
	a = BinarySearch(s, 10001)
	if a >= 0 {
		t.Error("should not found")
	}

	s = []int{1,2,3,4,5}
	a = BinarySearch(s, 4)
	if a != 3 {
		t.Errorf("BinarySearch got %d, want %d", a, 3)
	}
}
