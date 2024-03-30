package interview

import "sort"

// 在任意整数的数组里，输出所有满足 a+b+c=x*y 组合，要求 a, b, c, x, y不重复

func FindCombinations(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	mp := make(map[int][][3]int)
	res := [][]int{}
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[i-1] && i > 0 {
			continue
		}
		for j := 1; j < i; j++ {
			if nums[j] == nums[j-1] && j > 0 {
				continue
			}
			if nums[i] != nums[j] {
				for k := 0; k < j; k++ {
					if nums[i] != nums[k] && nums[j] != nums[k] {
						sum := nums[i] + nums[j] + nums[k]
						temp := [3]int{nums[i], nums[j], nums[k]}
						mp[sum] = append(mp[sum], temp)
					}
				}
			}
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] && i > 0 {
			continue
		}
		for j := 0; j < i; j++ {
			if nums[i] != nums[j] {
				temp := nums[i] * nums[j]
				if vals, ok := mp[temp]; ok {
					for m := 0; m < len(vals); m++ {
						tempNums := vals[m]
						tempmp := make(map[int]struct{})
						for k := 0; k < len(tempNums); k++ {
							tempmp[tempNums[k]] = struct{}{}
						}
						_, ok1 := tempmp[nums[i]]
						_, ok2 := tempmp[nums[j]]
						if !ok1 && !ok2 {
							restemp := []int{}
							for n := 0; n < len(vals[m]); n++ {
								restemp = append(restemp, vals[m][n])
							}
							restemp = append(restemp, nums[i])
							restemp = append(restemp, nums[j])
							res = append(res, restemp)
						}
						clear(tempmp)
					}
				}
			}
		}
	}
	return res
}
