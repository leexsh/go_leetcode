package leetcode

/*
题目：组合总和
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
只使用数字1到9
每个数字 最多使用一次
*/

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}

	var dfs func(path []int, index, sum int)
	dfs = func(path []int, index, sum int) {
		if sum == n && len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		if sum > k || len(path) > n {
			return
		}
		for i := index; i < 10; i++ {
			path = append(path, i)
			dfs(path, i+1, sum+i)
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 1, 0)
	return res
}
