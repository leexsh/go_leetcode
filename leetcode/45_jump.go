package leetcode

/*
题目：跳跃游戏
给你一个非负整数数组nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。


题解： https://leetcode.cn/problems/jump-game-ii/solution/tiao-yue-you-xi-ii-by-leetcode-solution/
https://leetcode.cn/problems/jump-game-ii/solution/45-tiao-yue-you-xi-iiliang-chong-tan-xin-q48k/
*/
func jump1(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j > i {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}
	return dp[len(nums)-1]
}
