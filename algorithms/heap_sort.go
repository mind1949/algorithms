package algorithms

// HeapSort 堆排序
func HeapSort(s []int) {
	if len(s) <= 1 {
		return
	}
	// 建堆
	buildMaxHeapify(s)
	// 排序
	for i := len(s) - 1; i > 0; i-- {
		s[0], s[i] = s[i], s[0]
		maxHeapify(s, 0, i-1)
	}
}

// buildMaxHeapify 建堆
func buildMaxHeapify(s []int) {
	h := len(s) - 1
	// 创建一个堆
	for i := (h - 1) / 2; i >= 0; i-- {
		maxHeapify(s, i, h)
	}
}

// maxHeapify 维持最大堆的性质
// 递归调用开销比循环控制结构开销大,
// 所以使用循环控制结构
func maxHeapify(s []int, l, h int) {
	root := l
	for {
		child := root*2 + 1
		if child > h {
			break
		}
		if child+1 <= h && s[child] < s[child+1] {
			child++
		}
		if s[root] >= s[child] {
			return
		}
		s[root], s[child] = s[child], s[root]
		root = child
	}
}
