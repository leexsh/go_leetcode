package leetcode

import "math"

//
//	mset := map[byte]int{}
//	n := len(s)
//
//	right, ans := -1, 0
//	for i := 0; i < n; i++ {
//		if i != 0 {
//			delete(mset, s[i-1])
//		}
//		for right+1 < n && mset[s[right+1]] == 0 {
//			right++
//			mset[s[right]]++
//		}
//		ans = max(ans, right-i+1)
//	}
//	return ans
// }
//
/*
题目：无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

*/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	i, j := 0, 0
	res := 0.0
	mp := make(map[byte]struct{})
	for j < len(s) {
		if _, ok := mp[s[j]]; !ok {
			mp[s[j]] = struct{}{}
			j++
			res = math.Max(res, float64(j-i))
		} else {
			delete(mp, s[i])
			i++
		}
	}
	return int(res)
}
