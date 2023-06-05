package leetcode

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

//
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/solution/dfsbian-li-golang-by-dz-lee/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mp := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var res []string
	var dfs func(int, string)
	dfs = func(i int, s string) {
		if i >= len(digits) {
			res = append(res, s)
			return
		}
		for _, ch := range mp[string(digits[i])] {
			dfs(i+1, s+string(ch))
		}
	}
	dfs(0, "")
	return res
}
