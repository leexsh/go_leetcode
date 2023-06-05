package leetcode

/*
题目：最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。
*/
// 中心扩展法
// https://leetcode.cn/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := expendFromCenter(s, i, i)
		l2, r2 := expendFromCenter(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

func expendFromCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
