package leetcode

/*
题目：全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
题解：

*/

func permute(nums []int) [][]int {
	var res = [][]int{}
	var dfs func(path []int)

	visited := make(map[int]struct{}, len(nums))
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, num := range nums {
			if _, ok := visited[num]; ok {
				continue
			}
			path = append(path, num)
			visited[num] = struct{}{}
			dfs(path)
			path = path[:len(path)-1]
			delete(visited, num)
		}
	}
	dfs([]int{})
	return res
}
