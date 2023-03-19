package leetcode

// https://leetcode.cn/problems/generate-parentheses/
// 回溯
// 括号生成
func generateParenthesis(n int) []string {
	var res []string
	if n == 0 {
		return res
	}
	var dfs func(str string, l, r int)
	dfs = func(str string, l, r int) {
		if l == 0 && r == 0 {
			res = append(res, str)
		}
		// 左括号比右括号多
		if l > r {
			return
		}
		if l > 0 {
			dfs(str+"(", l-1, r)
		}
		if r > 0 {
			dfs(str+")", l, r-1)
		}
	}
	dfs("", n, n)
	return res
}
