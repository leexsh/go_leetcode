package leetcode

import "sort"

/*
题目：字母的异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表
题解：
https://leetcode.cn/problems/group-anagrams/solution/zi-mu-yi-wei-ci-fen-zu-by-leetcode-solut-gyoc/
*/

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, strings := range mp {
		ans = append(ans, strings)
	}
	return ans
}
