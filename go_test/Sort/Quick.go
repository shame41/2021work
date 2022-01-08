package Sort

import (
	"go_test/Tool"
	"math/rand"
)

func QuickSort(data []int){
	quicksort(data, 0, len(data)-1)
}

func quicksort(data []int, lo, hi int){
	if hi <= lo + 10 {
		SelectSort(data)
		return
	}
	j := randPartition(data, lo, hi)
	//进行切分
	quicksort(data, lo, j-1)
	quicksort(data, j+1, hi)
}

func partition(data []int, lo,hi int) int{
	i := lo//左指针
	j := hi + 1//右指针
	v := data[lo]//我们取第一个元素为切分元素
	for{
		i++
		j--
		for data[i] < v {
			if i == hi {
				break
			}//从左到右扫描，找出第一个大于切分元素的数
			i++
		}
		for v < data[j] {
			if j == lo{
				break
			}//从右到左扫描，找出第一个小于切分元素的数
			j--
		}
		if i >= j{
			break
		}
		Tool.Swap(data, i, j)
	}
	//最后，把切分元素放回它应该在的位置
	Tool.Swap(data, lo, j)
	return j
}
func randPartition(data []int, lo,hi int) int{
	i := rand.Intn(hi - lo) + lo
	Tool.Swap(data, i, lo)
	return partition(data, lo, hi)
}