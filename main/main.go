package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

	for i := 0; i < 10; i++ {
		println(i)
	}
}
