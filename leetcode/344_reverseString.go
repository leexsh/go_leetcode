package leetcode

/*
题目：翻转字符串
*/
func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}
