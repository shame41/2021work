package algorithms

import "go_test/Tool"

// func Cut_Steel(p []int, n int) int{
// 	var profit []int
// 	for i := 0; i <= n; i++ {
// 		profit = append(profit, 0)
// 	}
// 	var q int
// 	for j := 1; j <= n; j++{
// 		q = -2147483648
// 		for i := 1; i <= j; i++ {
// 			q = Tool.Max(q, p[i] + profit[j - i])
// 		}
// 		profit[j] = q
// 	} 
// 	return profit[n]
// }

// func Cut_Steel(p []int, n int) int{
// 	if n == 0 {
// 		return 0
// 	}
// 	q := -2147483648
// 	for i := 1; i <= n ; i++ {
// 		q = Tool.Max(q, p[i] + Cut_Steel(p, n-1))
// 	}
// 	return q
// }
func Cut_Steel(p []int, n int) int{
	var profit []int
	for i := 0; i <= n; i++ {
		profit = append(profit, 0)
	}
	for i := 0; i <= n; i++ {
		profit[i] = -2147483648
	}
	return Memoized(p, profit, n)
}
func Memoized(p, r []int, n int) int{
	if r[n] >= 0 {
		return r[n]
	}
	var q int
	if n == 0 {
		q = 0
	}else{
		for i := 1; i <= n; i++ {
			q = Tool.Max(q, p[i] + Memoized(p, r, n-i))
		}
	}
	r[n] = q
	return q
}