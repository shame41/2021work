package Sort

import "go_test/Tool"

func ShellSort(data []int){
	N := len(data)
	h := 1;
	for h < N{
		h = 3*h + 1;
		//获得数列1,4,13,40,121,364,1093...
	}

	for h >= 1{
		//将数组变为每隔h就有序的数组
		for i := h; i < N; i++ {
			for j := i; j >= h && data[j] < data[j-1]; j=j-h {
				Tool.Swap(data, j, j-h)
			}
		}
		h = h/3//间隔会越来越小，直到h=1时，变为普通的插入排序
	}
}