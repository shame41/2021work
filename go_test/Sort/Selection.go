package Sort

import "go_test/Tool"

func SelectSort(data []int){
	N := len(data)
	for i := 0; i < N; i++{
		//将data[i]与data中最小的数交换
		min := i;//最小元素的索引
		for j := i + 1; j < N;  j ++{
			if data[j] < data[min]{
				min = j
			}
		}	
		Tool.Swap(data, min, i)
	}		
}