package Sort

import "go_test/Tool"

func InsertSort(data []int){
	N := len(data)
	for i := 1; i < N; i++ {
		//第一个元素已经排好
		//将未排序的元素通过冒泡的方法插入到已经排好的元素中
		for j := i; j > 0 && data[j]<data[j-1]; j-- {
			Tool.Swap(data,j, j-1)
		}
	}
}