package leetcode

/*
给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。

*/
import "math"

func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64

	for _, i := range nums {
		if i > a {
			a, b, c = i, a, b
		} else if i > b && a > i {
			b, c = i, b
		} else if i > c && i < b {
			c = i
		}
	}
	if c == math.MinInt64 {
		return 0
	}
	return c
}
