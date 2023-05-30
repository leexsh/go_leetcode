package main

import "fmt"

func longestConsecutive(nums []int) int {
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
			for _, ok := mp[curNum+1]; ok; {
				curNum++
				curLen++
			}
			if curLen > res {
				res = curLen
			}
		}
	}
	return res
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
