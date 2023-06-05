package leetcode

/*
题目：两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
*/

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for key, value := range nums {
		if idx, ok := mp[target-value]; ok {
			return []int{idx, key}
		}
		mp[value] = key
	}
	return nil
}
