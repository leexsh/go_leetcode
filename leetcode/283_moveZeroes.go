package leetcode

/*
题目：移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
*/
func moveZeroes(nums []int) {
	if len(nums) < 1 {
		return
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	for i := count; i < len(nums); i++ {
		nums[count] = 0
	}
}
