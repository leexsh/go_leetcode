package leetcode

/*
题目：多数元素
给定一个大小为 n 的数组nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
// 摩尔投票法
// 原理：在一个数组中一个元素出现的次数大于 n/2 = x,那也得必须要有x个元素才能抵消掉他，
// 但是2*x > n,因此剩下的元素一定是最多的那个。
func majorityElement(nums []int) int {
	res := 0
	cand := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == cand {
			res++
		} else if res == 0 {
			cand = nums[i]
		} else {
			res--
		}
	}
	return cand
}
