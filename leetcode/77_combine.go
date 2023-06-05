package leetcode

/*
题目：组合
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

题解：
*/
func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	res := [][]int{}
	var dfs func(begin int, path []int)
	dfs = func(begin int, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := begin; i <= n; i++ {
			path = append(path, i)
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{})
	return res
}
