package Tool

func MaxOfTwo(arr []int, a, b int) int{
	if len(arr) <= a {
		return b
	}else if len(arr) <= b{
		return a
	}else if arr[a] > arr[b]{
		return a
	}else{
		return b
	}
}
func MaxOfTwoF(arr []float64, a, b int) int{
	if len(arr) <= a {
		return b
	}else if len(arr) <= b{
		return a
	}else if arr[a] > arr[b]{
		return a
	}else{
		return b
	}
}