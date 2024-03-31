package leetcode

/*
给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。

示例 1：

输入：nums = [1,1,0,1,1,1]
输出：3
解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
示例 2:

输入：nums = [1,0,1,1,0,1]
输出：2
*/
func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, cnt := 0, 0
	for _, num := range nums {
		if num == 1 {
			cnt++
		} else {
			cnt = 0
		}
		res = max(res, cnt)
	}
	return res
}
