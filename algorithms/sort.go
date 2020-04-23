package algorithms

func IsSorted(slice []int) bool {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}
