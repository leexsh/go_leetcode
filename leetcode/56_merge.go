package leetcode

func merge56(nums [][]int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	
	})

	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(res) == 0 || res[len(res)-1][1] < nums[i][0] {
			res = append(res, nums[i])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], nums[i][1])
		}
	}
	return res

}