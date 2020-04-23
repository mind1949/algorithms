package algorithms

// QuickSort 快排
func QuickSort(s []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(s, p, r)
	QuickSort(s, p, q-1)
	QuickSort(s, q+1, r)
}

// partition 分区
func partition(s []int, p, r int) int {
	x := s[r]
	i := p - 1
	for j := p; j < r; j++ {
		if s[j] <= x {
			i++
			s[i], s[j] = s[j], s[i]
		}
	}
	s[i+1], s[r] = s[r], s[i+1]
	return i + 1
}
