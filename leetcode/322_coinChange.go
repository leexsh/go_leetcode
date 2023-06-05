package leetcode

/*
题目：零钱兑换
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
*/

func coinChange(coins []int, amount int) int {
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := amount + 1
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = max
	}
	dp[0] = 0
	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
