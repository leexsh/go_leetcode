package leetcode

import "sort"

/*
题目：子集
	给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

题解：
*/
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	sort.Ints(nums)
	var dfs func(begin int, path []int)
	dfs = func(begin int, path []int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		if len(path) > len(nums) {
			return
		}
		for i := begin; i < len(nums); i++ {
			if i > begin && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})
	return res
}
