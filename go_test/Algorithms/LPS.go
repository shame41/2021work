package algorithms

import "go_test/Tool"

func LPS(s string) int{
	n := len(s)
	if n == 1 {
		return 1
	}
	if n < 1 {
		return 0
	}
	var dp [][]int
	var maxDP int
	for i := 0; i < n; i++ {//dp[n][n]
		ar := make([]int, n)
		dp = append(dp, ar)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := i-1; j >= 0; j-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i-1][j+1] + 2
			}else{
				dp[i][j] = Tool.Max(dp[i][j+1], dp[i-1][j])
			}
			if dp[i][j] > maxDP {
				maxDP = dp[i][j]
			}
		}
	}
	return maxDP
}