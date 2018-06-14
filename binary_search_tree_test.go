package search

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree()
	if !b.IsEmpty() {
		t.Error("tree should be empty")
	}
	b.Insert(10)
	if b.IsEmpty() {
		t.Error("tree should not be empty")
	}
	b.Insert(11)
	b.Insert(9)
	b.Insert(8)
	b.Insert(12)
	if !b.Has(12) {
		t.Error("b.Has(12) return false, want true")
	}
	r, err := b.InorderTraverse()
	ln(r, cap(r), err)

	{
		b := NewBinarySearchTree()
		for i, j := 50, 50; i >= 0 && j < 100; {
			i--
			b.Insert(i)
			b.Insert(j)
			j++
		}
		r, err := b.InorderTraverse()
		ln(r, cap(r), err)
		for i := 0; i < len(r)-1; i++ {
			if r[i] >= r[i+1] {
				t.Error("sort result is wrong")
			}
		}
	}
}
