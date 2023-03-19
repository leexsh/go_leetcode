package leetcode

/*
题目：乘积最大子数组
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
*/

func maxProduct(nums []int) int {
	var max func(a, b int) int
	var min func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	res, imax, imin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			imin, imax = imax, imin
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])
		res = max(res, imax)
	}
	return res

}
