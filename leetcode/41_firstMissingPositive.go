package leetcode

/*
题目：缺失的第一个正数

描述：给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

题解：


*/
func firstMissingPositive(nums []int) int {
	mp := make(map[int]struct{}, len(nums))
	mx := 0
	for i := 0; i < len(nums); i++ {
		if i > mx {
			mx = nums[i]
		}
		mp[nums[i]] = struct{}{}
	}
	for i := 1; i < mx; i++ {
		if _, ok := mp[i]; !ok {
			return i
		}
	}
	return mx + 1
}
