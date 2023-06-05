package leetcode

import "math"

/*
题目：递增的三元子序列
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
*/
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt64
	for i := 1; i < n; i++ {
		if nums[i] > second {
			return true
		} else if nums[i] > first {
			second = nums[i]
		} else {
			first = nums[i]
		}
	}
	return false
}
