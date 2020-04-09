package sort

import (
	"testing"
)

var data = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestInsertionSort(t *testing.T) {
	ints := data[0:]
	InsertionSort(ints)
	if !IsSorted(ints) {
		t.Errorf("sorted %v", data[0:])
		t.Errorf("got %v", ints)
	}
}
