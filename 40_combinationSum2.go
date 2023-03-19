package leetcode

import "sort"

/*
题目：组合总和 II

描述：
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。

题解：见39

*/

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	var dfs func(tar, startPos int, path []int)
	dfs = func(tar, startPos int, path []int) {
		if tar == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := startPos; i < len(candidates); i++ {
			// 边界
			if candidates[i] > tar {
				return
			}
			// 去重
			if i > startPos && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(tar-candidates[i], i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(target, 0, []int{})
	return res
}
