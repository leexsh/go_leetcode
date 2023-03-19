package leetcode

import "math"

/*
题目：Pow(x, n)
题解：实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

*/
// todo
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	var num int = int(math.Abs(float64(n)))
	var res float64 = 1
	res = myPow(x, num>>1)
	res *= res
	if num&1 == 1 {
		res *= x
	}
	if n < 0 {
		return 1.0 / res
	}
	return res
}
