package leetcode

/*
26. 删除有序数组中的重复项
题目：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
题解：
*/

func removeDuplicates(nums []int) int {
	i, k := 0, 0
	for ; i < len(nums); i++ {
		if nums[k] != nums[i] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
