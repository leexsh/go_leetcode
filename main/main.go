package main

import (
	"fmt"
)

type ListNode struct {
	val  int
	next *ListNode
}
type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

/*
   快排
*/

func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	val := nums[0]
	l := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < val {
			l++
			nums[l], nums[i] = nums[i], nums[l]
		}
	}
	nums[0], nums[l] = nums[l], nums[0]
	QuickSort(nums[:l])
	QuickSort(nums[l+1:])
}

func main() {
	nums := []int{4, 25, 16, 7512, 121, 6465}
	QuickSort(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d\n", nums[i])
	}
}
