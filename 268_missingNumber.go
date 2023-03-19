package leetcode

/*
题目：丢失的数字
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

*/
func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
	}
	return sum - temp
}
