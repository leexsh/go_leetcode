package leetcode

/*
题目：存在重复元素
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
*/
func containsDuplicate(nums []int) bool {
	mp := make(map[int]int, len(nums))
	for _, num := range nums {
		mp[num]++
	}
	for _, num := range nums {
		if mp[num] > 1 {
			return true
		}
	}
	return false
}

func containsDuplicat1e(nums []int) bool {
	mp := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := mp[num]; ok {
			return true
		}
		mp[num] = struct{}{}
	}
	return false
}
