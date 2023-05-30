package leetcode

/*
题目：最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为O(n) 的算法解决此问题。

*/
func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	var maxCount int
	var tmpCount int
	for num := range m {
		if m[num-1] {
			continue
		}
		tmpCount = 0
		for m[num] {
			num++
			tmpCount++
		}
		if tmpCount > maxCount {
			maxCount = tmpCount
		}
	}
	return maxCount
}

func longestConsecutive1(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return res
	}
	mp := make(map[int]struct{})
	for _, num := range nums {
		mp[num] = struct{}{}
	}
	for i := range mp {
		if _, ok := mp[i-1]; !ok {
			curNum := i
			curLen := 1
			for {
				if _, ok := mp[curNum+1]; ok {
					curNum++
					curLen++
				} else {
					break
				}
			}
			if curLen > res {
				res = curLen
			}
		}
	}
	return res
}
