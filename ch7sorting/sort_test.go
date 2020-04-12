package sort

import (
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestInsertionSort(t *testing.T) {
	data := ints
	s := data[:]
	InsertionSort(s)
	if !IsSorted(s) {
		t.Errorf("sorted %v", ints[0:])
		t.Errorf("got %v", s)
	}
}

func TestHeapSort(t *testing.T) {
	data := ints
	s := data[:]
	HeapSort(s)
	if !IsSorted(s) {
		t.Errorf("sorted %v", ints[0:])
		t.Errorf("got %v", s)
	}
}
