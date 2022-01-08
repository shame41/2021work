package algorithms

func LCS_length(x, y []int) ([][]int, [][]int){
	m := len(x)
	n := len(y)

	var b [][]int
	var c [][]int
	for i := 0; i <= m; i++ {//b[m][n]
		ar := make([]int, n+1)
		b = append(b, ar)
	}
	for i := 0; i <= m; i++ {//c[m][n]
		ar := make([]int, n+1)
		c = append(c, ar)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				c[i][j] = c[i-1][j-1]+1
				b[i][j] = 1
			}else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = 2
			}else{
				c[i][j] = c[i][j-1]
				b[i][j] = 4
			}
		}
	}
	return c, b
}

func LCS_print(b [][]int, x []int, i, j int){
	if i == 0 || j == 0 {
		return
	}
	if b[i][j] == 1 {
		LCS_print(b, x, i-1, j-1)
		println(x[i-1])
	}else if b[i][j] == 2 {
		LCS_print(b, x, i-1, j)
	}else{
		LCS_print(b, x, i, j-1)
	}
}
func LCS(x,y []int){
	_, p := LCS_length(x, y)
	LCS_print(p, x, len(x), len(y))
}