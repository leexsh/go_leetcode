package leetcode

/*
题目：2的幂
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。
*/
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n&1 == 1 {
		return false
	}
	return isPowerOfTwo(n >> 1)
}

func isPowerOfTwo1(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}
