package leetcode

/*
题目：只出现一次的数字
给你一个整数数组nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法且不使用额外空间来解决此问题。

*/
func singleNumber2(nums []int) int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	for key, val := range mp {
		if val == 1 {
			return key
		}
	}
	return 0
}
