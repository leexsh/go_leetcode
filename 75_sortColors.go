package leetcode

/*
题目：颜色分类
给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
必须在不使用库的sort函数的情况下解决这个问题。
题解：
荷兰国旗
*/
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	cur := 0
	for cur <= right {
		if nums[cur] < 1 {
			nums[cur], nums[left] = nums[left], nums[cur]
			left++
			cur++
		} else if nums[cur] == 1 {
			cur++
		} else {
			nums[cur], nums[right] = nums[right], nums[cur]
			right--
		}
	}
}
