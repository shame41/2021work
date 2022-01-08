package algorithms

func LIS(a []int) int{
	n := len(a)
	dp := make([]int, n)
	dpMax := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if a[j] < a[i] {
				dp[i] = dp[j] + 1
				if dp[j] >= dpMax {
					dpMax = dp[i]
				}
			}
		}
	}
	return dpMax
}