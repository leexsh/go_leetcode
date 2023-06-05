package leetcode

/*
题目：最大子数组之和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
题解：
*/
func maxSubArray(nums []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res, dp := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		dp = max(dp+nums[i], nums[i])
		res = max(dp, res)

	}
	return res
}
