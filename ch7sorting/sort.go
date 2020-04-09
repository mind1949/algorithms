package sort

func InsertionSort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

func IsSorted(slice []int) bool {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}
