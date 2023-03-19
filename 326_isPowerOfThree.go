package leetcode

/*
题目：3的幂
*/
func isPowerOfThree(n int) bool {
	if n == 0 {
		return true
	}
	if n == 1 {
		return true
	}
	if n%3 == 0 {
		n /= 3
		return isPowerOfThree(n)
	}
	return false
}
