package leetcode

/*
题目：长度最小的子树组
给定一个含有 n 个正整数的数组和一个正整数 target 。

*/
func minSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1
	i, j := 0, 0
	sum := 0
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for j < len(nums) {
		sum += nums[j]
		for sum >= target {
			res = min(res, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
