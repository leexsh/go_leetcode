package leetcode

import "sort"

/*
题目： 全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
题解：https://leetcode.cn/problems/permutations-ii/
*/

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	sort.Ints(nums)
	helper([]int{}, nums, used, &res)
	return res
}

func helper(path, nums []int, used []bool, res *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i-1 >= 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		helper(path, nums, used, res)
		path = path[0 : len(path)-1]
		used[i] = false
	}
}
