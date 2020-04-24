package algorithms

import (
	"math/rand"
	"time"
)

func IsSorted(slice []int) bool {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}

func Random(p, r int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(r-p+1) + p
}
