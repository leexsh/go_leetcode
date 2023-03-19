package leetcode

/*
题目：删除重复的元素
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

*/

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slow, num := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			num++
		} else {
			num = 1
		}
		if num <= 2 {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}
