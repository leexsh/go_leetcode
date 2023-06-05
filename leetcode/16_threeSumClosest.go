package leetcode

import (
	"math"
	"sort"
)

/*
题目：最接近的三数之和
给你一个长度为 n 的整数数组nums和 一个目标值target。请你从 nums 中选出三个整数，使它们的和与target最接近。

返回这三个数的和。

*/
// 见cpp版本

func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		L, R := i+1, len(nums)-1
		for L < R {
			temp := nums[i] + nums[R] + nums[L]
			if math.Abs(float64(temp-target)) < math.Abs(float64(res-target)) {
				res = temp
			}
			if temp-target == 0 {
				return temp
			} else if temp-target > 0 {
				R--
			} else {
				L++
			}
		}
	}
	return res
}
