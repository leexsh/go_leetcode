package leetcode

import "math"

/*
题目：平方数之和
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。
*/
func judgeSquareSum(c int) bool {
	a := 0
	b := int(math.Sqrt(float64(c)))
	for a <= b {
		mid := int(a*a + int(b*b))
		if mid == c {
			return true
		} else if mid < c {
			a++
		} else {
			b--
		}
	}
	return false
}
