package Sort

import "go_test/Tool"
func RadixSort(a []int){
	min , max := 0, 0
	for _, i := range(a){
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}//得到最大数
	if min < 0 {
		for i := range(a){
			a[i] -= min
		}
	}
	for i := 1; max > 0; i++{
		Counting(a, i)//对每一位进行计数排序
		max = max/10
	}
	if min < 0 {
		for i := range(a){
			a[i] += min
		}
	}
}

func Counting(a []int,k int){
	var c []int
	var max int
	b := make([]int, len(a))
	if k == 0 {//普通计数排序
		for _, i := range(a){
			if i > max {
				max = i
			}
		}//得到最大数
		for i := 0; i <= max; i++{
			c = append(c, 0)
		}//初始化计数数组
		for j := 0; j < len(a); j++ {
			c[a[j]] = c[a[j]] + 1
		}//对数组a中出现的数组计数
		for i := 1; i <= max; i++ {
			c[i] = c[i] + c[i-1]
		}//将计数数组转换为位移数组，索引上的值就是索引在数组a中应在的位置

		for j := len(a) - 1; j >= 0; j--{
			b[c[a[j]] - 1] = a[j]
			c[a[j]] = c[a[j]] - 1
		}
	}else{//按十位计数排序
		for _, i := range(a){
			temp := Tool.Digit(i, k)
			if temp > max {
				max = temp
			}
		}//得到最大数
		for i := 0; i <= max; i++{
			c = append(c, 0)
		}//初始化计数数组
		for j := 0; j < len(a); j++ {
			temp := Tool.Digit(a[j], k)
			c[temp] = c[temp] + 1
		}//对数组a中出现的数组计数
		for i := 1; i <= max; i++ {
			c[i] = c[i] + c[i-1]
		}//将计数数组转换为位移数组，索引上的值就是索引在数组a中应在的位置
		for j := len(a) - 1; j >= 0; j--{
			temp := Tool.Digit(a[j], k)
			b[c[temp] - 1] = a[j]
			c[temp] = c[temp] - 1
		}
	}
	for i:= range(a){
		a[i] = b[i]
	}
}

func LocalCounting(a []int){//算法导论8-2-e，原址计数排序
	var c []int
	var max int
	for _, i := range(a){
		if i > max {
			max = i
		}
	}//得到最大数
	for i := 0; i <= max; i++{
		c = append(c, 0)
	}//初始化计数数组
	for j := 0; j < len(a); j++ {
		c[a[j]] = c[a[j]] + 1
	}//对数组a中出现的数组计数
	for i := 1; i <= max; i++ {
		c[i] = c[i] + c[i-1]
	}//将计数数组转换为位移数组，索引上的值就是索引在数组a中应在的位置
	for i := len(a)-1; i >= 0; i--{
		p := c[a[i]] - 1
		for p > i{
			c[a[i]] = c[a[i]] - 1
			Tool.Swap(a, i, p)
			p = c[a[i]] - 1
		} 
	}
}