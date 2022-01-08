package Tool

import "math"

// 数字，右数第几位，从1开始
func Digit(num int, loc int) int {
	x := num % int(math.Pow10(loc)) / int(math.Pow10(loc-1))
    return x
}