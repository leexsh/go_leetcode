package leetcode

import "sort"

/*
题目：三数之和
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。
*/
// https://leetcode.cn/problems/3sum/solution/suan-fa-si-wei-yang-cheng-ji-er-fen-cha-5bk43/

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	right := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L, R := i+1, right
		for L < R {
			sum := nums[i] + nums[R] + nums[L]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[R], nums[L]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if sum > 0 {
				R--
			} else {
				L++
			}
		}
	}
	return res
}
