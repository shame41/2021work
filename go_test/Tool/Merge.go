package Tool

func Merge(data []int, lo, mid, hi int){
	i := lo
	j := mid + 1
	aux := make([]int, len(data))
	for k := lo; k <= hi; k++ {
	//将data[lo..hi]整个复制到aux[lo..hi]	
		aux[k] = data[k]
	}
	for k := lo; k <= hi; k++ {
		if i > mid{//左半边已经没有元素的情况
			data[k] = aux[j]
			j++
		}else if j > hi {//右半边已经没有元素的情况
			data[k] = aux[i]
			i++
		}else if aux[j] < aux[i] {//左半边当前值大于右半边当前值
			data[k] = aux[j]
			j++
		}else{
			data[k] = aux[i]
			i++
		}
	}
}