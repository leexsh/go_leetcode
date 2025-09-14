package algorithm

/*
   快排
*/

func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	value := nums[0]
	cur := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < value {
			cur++
			nums[cur], nums[i] = nums[i], nums[cur]
		}
	}
	nums[cur], nums[0] = nums[0], nums[cur]
	QuickSort(nums[:cur])
	QuickSort(nums[cur+1:])
}
