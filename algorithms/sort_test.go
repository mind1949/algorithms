package algorithms

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

func TestSelectionSort(t *testing.T) {
	data := ints
	s := data[:]
	SelectionSort(s)
	if !IsSorted(s) {
		t.Errorf("sorted %v", ints[0:])
		t.Errorf("got %v", s)
	}
}

func TestMerge(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	m := len(data) / 2
	merge(data[:m], data[m:])
	expect := []int{1, 2, 3, 4, 5, 6}
	for i, v := range expect {
		if data[i] != v {
			t.Errorf("expect %v", expect)
			t.Errorf("got %v", data)
			return
		}
	}
}

func TestMergeSort(t *testing.T) {
	data := ints
	s := data[:]
	MergeSort(s)
	if !IsSorted(s) {
		t.Errorf("sorted %v", ints[0:])
		t.Errorf("got %v", s)
	}
}
