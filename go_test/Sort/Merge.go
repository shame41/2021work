package Sort

import (
	"go_test/Tool"
	"math"
)

func MergeUD(data []int){//自顶向下归并排序
	mergesort(data, 0, len(data)-1)
}

func MergeDU(data []int){
	N := len(data)
	for sz:=1; sz<N; sz=sz+sz {//sz是每个子数组的大小
		for lo := 0; lo < N-sz; lo=lo+sz+sz {
			Tool.Merge(data, lo, lo+sz-1,int(math.Min(float64(lo+sz+sz-1),float64(N-1))))
			//值得一提的是，math.Min方法的参数只能是float64
		}
	}
}

func mergesort(data []int, lo, hi int){
	//排序数组a[lo..hi]排序
	if hi <= lo {
		return
	}
	mid := lo + (hi - lo)/2
	mergesort(data, lo, mid)//排序左半边
	mergesort(data, mid+1, hi)//排序右半边
	Tool.Merge(data, lo, mid, hi)//归并结果
}
