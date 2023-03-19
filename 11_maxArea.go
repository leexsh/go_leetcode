package leetcode

// 题目: https://leetcode.cn/problems/container-with-most-water/
/*
题目：盛最多水的容器
给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
*/
func maxArea(height []int) int {
	i, j, max := 0, len(height)-1, 0
	for i < j {
		var temp int = 0
		if height[i] > height[j] {
			temp = height[j] * (j - i)
			j--
		} else {
			temp = height[i] * (j - i)
			i++
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
