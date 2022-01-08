package Sort

import (
	"go_test/Tool"
	"math/rand"
)

func Sort3Quick(data []int){
	sort3way(data, 0, len(data)-1)
}

func sort3way(data []int, lo, hi int){
	//三切分快速排序，用于处理有许多相同数据的数组的排序
	if hi <= lo + 10 {
		SelectSort(data)//对于小数组，使用选择排序
		return
	}
	lt := lo//小于切分元素的集合的边界
	i := lo + 1//当前元素
	gt := hi//大于切分元素的集合的边界
	key := rand.Intn(hi - lo) + lo//随机选择驻点
	Tool.Swap(data, key, lo)
	v := data[lo]//切分元素
	for i <= gt {
		if data[i] < v {//放入小于切分元素的集合里
			Tool.Swap(data, lt, i)//注意，这里不是和切分元素交换，而是与边界元素交换
			lt++//让集合变大
			i++
		}else if data[i] > v{//放入大于切分元素的集合里
			Tool.Swap(data, i, gt)
			gt--//让集合变大
		}else{//此元素保留在本地，本地就是与切分元素相等的集合
			i++
		}
	}
	sort3way(data, lo, lt-1)
	sort3way(data, gt+1, hi)
}