package sort

// MergeSort 归并排序
func MergeSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}

	// 分解
	m := n / 2
	l := s[:m]
	r := s[m:]
	// 解决
	MergeSort(l)
	MergeSort(r)
	// 合并
	merge(l, r)
}

// merge 合并
func merge(l, r []int) {
	i := 0
	j := 0
	tmp := make([]int, len(l)+len(r))
	for k := 0; k < len(tmp); k++ {
		if i == len(l) {
			copy(tmp[k:], r[j:])
			break
		}
		if j == len(r) {
			copy(tmp[k:], l[i:])
			break
		}
		if l[i] <= r[j] {
			tmp[k] = l[i]
			i++
			continue
		}
		tmp[k] = r[j]
		j++
	}
	copy(l, tmp[:len(l)])
	copy(r, tmp[len(l):])
}
