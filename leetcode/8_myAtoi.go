package leetcode

import "math"

/*
题目：atoi
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
*/
// https://leetcode.cn/problems/string-to-integer-atoi/solution/go-0-ms-zai-suo-you-go-ti-jiao-zhong-ji-bvpku/
func myAtoi(s string) int {
	res, i, len := 0, 0, len(s)
	flag := 1

	for i < len && s[i] == ' ' {
		i++
	}
	if i < len {
		if s[i] == '-' {
			flag = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}
	for i < len && s[i] >= '0' && s[i] <= '9' {
		res = res*10 + int(s[i]-'0')
		if res*flag > math.MaxInt32 {
			return math.MaxInt32
		} else if res*flag < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	return flag * res
}
