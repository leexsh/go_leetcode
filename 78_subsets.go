package leetcode

/*
题目：子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

题解：
*/
func subsets(nums []int) [][]int {
	res := [][]int{}
	var dfs func(begin int, path []int)
	dfs = func(begin int, path []int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		if begin > len(nums) {
			return
		}
		for i := begin; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{})
	return res
}
