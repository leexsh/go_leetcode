package leetcode

/*
给你一个字符串 s，最多 可以从中删除一个字符。

请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。



示例 1：

输入：s = "aba"
输出：true
示例 2：

输入：s = "abca"
输出：true
解释：你可以删除字符 'c' 。
示例 3：

输入：s = "abc"
输出：false

*/

func isPalindromeStr(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return isPalindromeStr(s, i+1, j) || isPalindromeStr(s, i, j-1)
		}
		i++
		j--
	}
	return true
}
