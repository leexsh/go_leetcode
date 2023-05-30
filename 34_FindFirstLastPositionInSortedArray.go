package leetcode

/*
34. 在排序数组中查找元素的第一个和最后一个位置
题目：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
题解：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/solution/er-fen-cha-zhao-hou-bian-jie-kuo-zhan-by-bbrw/
*/
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			left, right = mid, mid
			for left > 0 && nums[left] == nums[left-1] {
				left--
			}
			for right < len(nums)-1 && nums[right] == nums[right+1] {
				right++
			}
			return []int{left, right}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}
