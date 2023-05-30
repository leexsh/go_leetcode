package leetcode

/*
题目：存在重复的元素
给你一个整数数组nums 和一个整数k ，判断数组中是否存在两个 不同的索引i和j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。

*/

// map
func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, ok := mp[num]; ok && (i-index) <= k {
			return true
		}
		mp[num] = i
	}
	return false
}

// 滑动窗口
func containsNearbyDuplicate1(nums []int, k int) bool {
	set := make(map[int]struct{}, len(nums))
	for i, num := range nums {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
