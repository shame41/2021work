package Sort

func FindMinK(a []int, k int) int{
	return findMinK(a, 0, len(a)-1, k)
}

func findMinK(a []int, lo, hi, k int) int{
	if lo == hi {
		return a[lo]
	}

	q := randPartition(a, lo, hi)//切分元素的位置
	distance := q - lo + 1//切分元素与lo的相对距离
	if k == distance {//如果切分元素到最左边的距离是k，那么切分元素就是第K小的元素
		return a[q]//所要找的顺序统计量就是切分元素
	}else if k < distance{
		return findMinK(a, lo, q-1, k)
	}else{
		return findMinK(a, q+1, hi, k - distance)
	}
}