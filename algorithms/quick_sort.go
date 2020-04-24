package algorithms

// QuickSort 快排
func QuickSort(s []int, p, r int) {
	if p >= r {
		return
	}
	q := randomizedPartition(s, p, r)
	QuickSort(s, p, q-1)
	QuickSort(s, q+1, r)
}

// randomizedPartition 随机分区
// 通过随机抽样似的对于任何输入都能获得较好的期望性能
func randomizedPartition(s []int, p, r int) (q int) {
	i := Random(p, r)
	s[i], s[r] = s[r], s[i]
	return partition(s, p, r)
}

// partition 分区
// 经过分区处理后
// s[p:q]中任一元素小于或等于s[q]
// s[q+1:r]中任一元素大于s[q]
func partition(s []int, p, r int) (q int) {
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

/*
证明partition函数的正确性
* 循环不变式:
s[p:i]中任一元素≤主元x, s[i:j]中任一元素≥主元,
*/
