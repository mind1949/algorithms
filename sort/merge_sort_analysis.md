# 证明/分析归并排序

## 代码
```go
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
```

## 证明

### 证明`merge`算法正确性
* 循环不变式:
>tmp[:k-1]由l与r中最小的k-0个元素按顺序排序组成, l[i:]中的元素l[i]为l[i:]中最小的元素,
>r[j:]中的元素r[j]为r[j:]中最小的元素
* 初始化:
>在初始化时k=0, m[:0]中的元素个数为零, 由l与r中最小的0-0个元素组成;
>i = 0, l[0:]中的l[0]为l[0:]中最小的元素;
>j = 0, r[0:]中国的r[0]为r[0:]中的最下元素.
>此时驯化不变式为真
* 保持:
>假设在某一次迭代前, tmp[:k-1]由l与r中最小的k-0个元素按顺序排序组成, l[i:]中的元素l[i]为l[i:]中最小的元素,
r[j:]中的元素r[j]为r[j:]中最小的元素, 不变式为真, 则经过迭代后tmp[:k], tmp[k]为l[i:]与r[j:]中的最小元素组成, 此时m[:k]由l与r中最小的k-0+1个元素组成, 若l[i] <= r[j], 则l[i+1:]中的l[i+1]为l[i+1:]中的最小值, r[j]为r[j:]中的最小值, 所以在下次循环迭代之前, 循环不变式依然为真.
同理可得当l[i] > r[j]时, 依然满足循环不变式为真.
* 终止:
> 当迭代终止时
k = len(tmp), 等于len(l)+len(r), 根据循环不变式为真, 有: tmp[:len(l)+len(r)]的元素由l与r中len(l)+len(r)个元素按顺序排序组成. 此时可以证明算法正确性.

### 证明`MergeSort`算法正确性


## 分析

### 分析`merge`算法时间复杂度

### 分析`MergeSort`算法时间复杂度