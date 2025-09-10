package algorithm

/*
   快排
*/

func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	val := nums[0]
	l, r := 0, len(nums)-1
	for i := 0; i < r; {
		if nums[i] < val {
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		} else if nums[i] > val {
			nums[r], nums[i] = nums[i], nums[r]
			r--
		} else {
			i++
		}
	}
	QuickSort(nums[:l])
	QuickSort(nums[r+1:])
}