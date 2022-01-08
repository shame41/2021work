package Sort

func LinearMaxSub(arr []int) (int, int, int){
	var s1,s2,e1,e2 int
	s1, e1, s2, e2 = 0, 0, 0, 0
	max := 0
	sum := 0
	flag := 0
	for i := 0; i < len(arr); i++{
		sum += arr[i]
		if sum > 0 && flag == 0{
			e1 = i
		}else if sum > 0 && flag != 0 {
			e2 = i
		}
		if sum > max {
			flag = 0
			s1 = s2
			e1 = e2
			max = sum
		}
		if sum < 0 {
			sum = 0
			flag = 1
			s2 = i+1
			e2 = s2
		}
	} 
	return s1, e1, max
}