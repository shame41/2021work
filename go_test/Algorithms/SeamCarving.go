package algorithms


func SeamCarving(d [][]int) []int{
	n := len(d[0])
	m := len(d)
	min := make([]int, m)
	var minIndex int
	truesum := 9223372036854775807
	flag := make([]int, m)
	for key := 0; key < n; key++ {
		minIndex = key
		for j := 1; j < m; j++ {
			min[j], minIndex = MinOfThree(d[j],minIndex)
		}
		sum := d[0][key]
		for i := 0; i < m; i++ {
			sum = sum + min[i]
		}
		if sum < truesum{
			truesum = sum
			for i := range(min) {
				flag[i] = min[i]
			}
			flag[0] = d[0][key]
		}
	}
	return flag
}

func MinOfThree(a []int, k int) (int, int){
	if k == 0 {
		if a[k] < a[k+1] {
			return a[k], k
		}else{
			return a[k+1], k+1
		}
	}
	if k == len(a)-1 {
		if a[k] > a[k-1] {
			return a[k-1], k-1
		}else{
			return a[k], k
		}
	}
	min := 2147483647
	var minIndex int
	for i := k-1; i <= k+1; i++ {
		if a[i] < min {
			min = a[i]
			minIndex = i
		}
	}
	return min, minIndex
}