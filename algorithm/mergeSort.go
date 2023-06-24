package algorithm

/*
归并排序
*/

func mergeSort(nums []int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	l := mergeSort(nums[:len(nums)/2])
	r := mergeSort(nums[len(nums)/2:])
	return merge(l, r)
}

func merge(l, r []int) []int {
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(l) && j <= len(r) {
		if l[i] < r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	res = append(res, l[i:]...)
	res = append(res, r[j:]...)
	return res
}
