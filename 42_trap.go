package leetcode

/*
题目：接雨水
https://leetcode.cn/problems/trapping-rain-water/
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
题解：https://leetcode.cn/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode-solution-tuvc/
*/

// 双指针
func trap(height []int) int {
	left, right := 0, len(height)-1
	lMax, rMax := height[left], height[right]
	ans := 0
	for left < right {
		lMax = max(lMax, height[left])
		rMax = max(rMax, height[right])
		if height[left] < height[right] {
			ans += lMax - height[left]
			left++
		} else {
			ans += rMax - height[right]
			right--
		}
	}
	return ans
}
