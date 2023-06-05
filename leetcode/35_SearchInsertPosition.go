package leetcode

/*
35.搜索插入位置
题目：https://leetcode.cn/problems/search-insert-position/
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

题解：
*/

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	res := len(nums)
	for left <= right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
