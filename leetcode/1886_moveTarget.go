package leetcode

/*
给定一个数组 nums 以及一个整数 target 。
你需要把数组中等于target的元素移动到数组的最前面，并且其余的元素相对顺序不变。
你的所有移动操作都应该在原数组上面操作。

*/

func MoveTarget(nums []int, target int) {
	// write your code here
	if len(nums) == 0 {
		return
	}
	count := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			count++
		} else {
			if count > 0 {
				nums[i+count] = nums[i]
			}
		}
	}
	for i := 0; i < count; i++ {
		nums[i] = target
	}

}
