package leetcode

// 题解：https://leetcode.cn/problems/roman-to-integer/solution/luo-ma-shu-zi-zhuan-zheng-shu-by-leetcod-w55p/
/*
题目：罗马数字转整数
https://leetcode.cn/problems/roman-to-integer/

*/
var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		val := symbolValues[s[i]]
		if i < n-1 && val < symbolValues[s[i+1]] {
			ans -= val
		} else {
			ans += val
		}
	}
	return
}
