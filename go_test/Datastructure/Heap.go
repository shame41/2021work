package datastructure

import "go_test/Tool"

type HeapForInt []int

func (a HeapForInt) swim(k int){
	for k > 1 {
		if a[k] < a[k/2]{
			Tool.Swap(a, k, k/2)
		}else{
			break
		}
	}
}

func (a *HeapForInt) sink(k int){
	size := len(*a) - 1
	for k <= size/2 {
		key := Tool.MaxOfTwo(*a, 2*k, 2*k + 1)
		if (*a)[k] < (*a)[key]{
			Tool.Swap(*a, k, key)
			k = 2*k
		}else{
			break
		}
	}
}

func (a *HeapForInt) Max() int{
	max := (*a)[1]
	size := len(*a)-1
	if size == 1 {
		return max
	}
	Tool.Swap(*a, 1, size)
	*a = (*a)[:size]
	a.sink(1)
	return max
}

func (a *HeapForInt) Build(arr []int) {
	for _, k := range(arr){
		*a = append(*a, k)
	}
	for i := len(arr)/2 + 1; i > 0; i-- {
		a.sink(i)
	}
}

func (a *HeapForInt) Add(k int){
	*a = append(*a, k)
	a.swim(len(*a))
}

type HeapForFloat []float64

func (a *HeapForFloat) swim(k int){
	for k > 1 {
		if (*a)[k] < (*a)[k/2]{
			Tool.SwapF(*a, k, k/2)
		}else{
			break
		}
	}
}

func (a *HeapForFloat) sink(k int){
	size := len(*a) - 1
	for k <= size/2 {
		key := Tool.MaxOfTwoF(*a, 2*k, 2*k + 1)
		if (*a)[k] < (*a)[key]{
			Tool.SwapF(*a, k, key)
			k = 2*k
		}else{
			break
		}
	}
}

func (a *HeapForFloat) Max() float64{
	max := (*a)[1]
	size := len(*a)
	Tool.SwapF(*a, 1, size - 1)
	*a = (*a)[:size - 1]
	a.sink(1)
	return max
}

func (a *HeapForFloat) Build() {
	*a = append(*a, 0)
}

func (a *HeapForFloat) Add(k float64){
	*a = append(*a, k)
	a.swim(len(*a))
}
