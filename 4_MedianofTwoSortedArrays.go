package leetcode

/*
题目：寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2, m, n := len(nums1), len(nums2), 0, 0
	// var res []int
	res := make([]int, 0, 0)
	for m < l1 && n < l2 {
		if nums1[m] < nums2[n] {
			res = append(res, nums1[m])
			m++
		} else {
			res = append(res, nums2[n])
			n++
		}
	}
	res = append(res, nums1[m:l1]...)
	res = append(res, nums2[n:l2]...)
	len := l1 + l2
	if len%2 == 1 {
		return float64(res[len/2])
	}
	return float64(res[len/2]+res[len/2-1]) / 2.0
}
