package Sort

import datastructure "go_test/Datastructure"

func HeapSort(arr []int){
	var heap datastructure.HeapForInt
	//println(heap.Max())
	heap.Build(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = heap.Max()
	}
}