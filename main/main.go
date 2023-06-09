package main

import (
	"fmt"
)

func mergeTwoArray(nums1, nums2 []int) []int {
	l1, l2, l := len(nums1)-1, len(nums2)-1, 0
	res := make([]int, l1+l2)
	i, j := 0, 0
	for i <= l1 && j <= l2 {
		if nums1[i] < nums2[j] {
			res[l] = nums1[i]
			l++
			i++
			for i <= l1 && l > 0 {
				if nums1[i] == res[l-1] {
					i++
				} else {
					break
				}
			}
		} else if nums1[i] == nums2[j] {
			res[l] = nums1[i]
			i++
			j++
			l++
			for i <= l1 && l > 0 {
				if nums1[i] == res[l-1] {
					i++
				} else {
					break
				}
			}
			for j <= l2 && l > 0 {
				if nums2[j] == res[l-1] {
					j++
				} else {
					break
				}
			}
		} else {
			res[l] = nums2[j]
			j++
			l++
			for j <= l2 && l > 0 {
				if nums2[j] == res[l-1] {
					j++
				} else {
					break
				}
			}
		}
	}

	for i <= l1 {
		res[l] = nums1[i]
		l++
		i++
		for i <= l1 && l > 0 {
			if nums1[i] == res[l-1] {
				i++
			} else {
				break
			}
		}
	}
	for j <= l2 {
		res[l] = nums2[j]
		l++
		j++
		for j <= l2 && l > 0 {
			if nums2[j] == res[l-1] {
				j++
			} else {
				break
			}
		}
	}
	return res[:l]
}

func main() {
	n1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3}
	n2 := []int{0, 0, 2, 2, 2, 3, 3, 5, 5, 5}
	res := mergeTwoArray(n1, n2)
	for _, re := range res {
		fmt.Println(re)

	}
}
