package leetcode

/*
题目：打家劫舍2
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
*/

func max_rob(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob2_interal(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max_rob(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max_rob(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max_rob(nums[0], nums[1])
	}
	return max(rob2_interal(nums[1:]), rob2_interal(nums[:len(nums)-1]))
}
