package leetcode

/*
题目：整数拆分
*/
func integerBreak(n int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, n+1)
	dp[1] = 0
	dp[0] = 0
	dp[2] = 1
	for i := 3; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			temp := max(j*(i-j), j*dp[i-j])
			curMax = max(curMax, temp)
		}
		dp[i] = curMax
	}
	return dp[n]
}
