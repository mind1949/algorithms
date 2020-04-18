package sort

// SelectionSort 选择排序
func SelectionSort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if s[j] < s[minIdx] {
				minIdx = j
			}
		}
		s[i], s[minIdx] = s[minIdx], s[i]
	}
}
