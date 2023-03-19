package leetcode

/*
题目：回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
*/
// https://leetcode.cn/problems/palindrome-number/solution/tu-jie-guan-fang-tui-jian-ti-jie-hui-wen-pqvq/

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	res := 0
	for x > res {
		num := x % 10
		res = res*10 + num
		x /= 10
	}
	return x == res || res/10 == x

}
