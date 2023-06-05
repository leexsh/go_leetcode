package leetcode

/*
题目：跳跃游戏
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。

题解：

*/

func canJump(nums []int) bool {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	mixDis := 0
	for i := 0; i < len(nums); i++ {
		if i <= mixDis {
			mixDis = max(mixDis, i+nums[i])
			if mixDis >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
