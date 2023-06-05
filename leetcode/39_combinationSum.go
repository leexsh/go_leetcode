package leetcode

import "sort"

/*
题目：39. 组合总和

描述：给你一个 无重复元素 的整数数组candidates 和一个目标整数target，找出candidates中可以使数字和为目标数target 的 所有不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

题解：https://leetcode.cn/problems/combination-sum/solution/shou-hua-tu-jie-zu-he-zong-he-combination-sum-by-x/

22-24为什么要重新拷贝一个切片？
	切片是一个结构体, 主要有指针(指向底层数组)和len和cap字段组成, 如果我们直接用temp加入到结果集,append函数帮我们拷贝一份过去,
	但是 !!! 指针也同样拷贝进去了, 所以指向的是同一个底层数组 后续我们对temp进行更改结果集中的结果切片也会进行更改,
	因为他两的指针是指向同一个底层数组的, 所以我们要新开辟一个切片把结果拷贝过去, 这样我们后续对于temp切片的改变才不会影响到res中的结果集
*/

func combinationSum(candidates []int, target int) [][]int {
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
			if candidates[i] > tar {
				break
			}
			path = append(path, candidates[i])
			dfs(tar-candidates[i], i, path)
			path = path[:len(path)-1]
		}
	}
	dfs(target, 0, []int{})
	return res
}
