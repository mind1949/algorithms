package sort

// HeapSort 堆排序
func HeapSort(s []int) {
	if len(s) <= 1 {
		return
	}

	h := len(s) - 1
	// 创建一个堆
	for i := (h - 1) / 2; i >= 0; i-- {
		heaptify(s, i, h)
	}
	// 排序
	for i := h; i > 0; i-- {
		s[0], s[i] = s[i], s[0]
		heaptify(s, 0, i-1)
	}
}

func heaptify(s []int, l, h int) {
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
