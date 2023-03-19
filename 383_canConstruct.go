package leetcode

/*
题目：赎金信
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

*/
func canConstruct(ransomNote string, magazine string) bool {
	mp := map[byte]int{}
	for _, ch := range magazine {
		mp[byte(ch-'a')]++
	}
	for _, ch := range ransomNote {
		mp[byte(ch-'a')]--
		if mp[byte(ch-'a')] < 0 {
			return false
		}
	}
	return true
}
