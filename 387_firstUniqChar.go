package leetcode

/*
题目：字符串中的第一个唯一字符
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
*/

func firstUniqChar(s string) int {
	mp := map[byte]int{}
	for _, ch := range s {
		mp[byte(ch-'a')]++
	}
	for i, ch := range s {
		if mp[byte(ch-'a')] == 1 {
			return i
		}
	}
	return -1
}
