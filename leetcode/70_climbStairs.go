package leetcode

/*
题目： 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢

题解：
*/

func climbStairs(n int) int {
	a, b, j := 1, 2, 0
	if n <= 2 {
		return n
	}
	for i := 3; i <= n; i++ {
		j = a + b
		a, b = b, j
	}
	return j
}
