package leetcode

import (
	"math"
)

/*
题目: 整数翻转
*/
// 整数翻转
// https://leetcode.cn/problems/reverse-integer/solution/zheng-shu-fan-zhuan-by-leetcode-solution-bccn/
func reverse(x int) int {
	res := 0
	for x != 0 {
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}
		num := x % 10
		x /= 10
		res = res*10 + num
	}
	return res
}
