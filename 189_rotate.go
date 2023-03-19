package leetcode

/*
题目：轮转数组
给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
*/
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse_1(nums)
	reverse_1(nums[:k])
	reverse_1(nums[k:])
}

func reverse_1(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}
