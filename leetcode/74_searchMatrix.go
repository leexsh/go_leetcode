package leetcode

/*
题目：搜索二维矩阵
编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：
每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
题解：
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix)-1, 0
	for n < len(matrix[0]) && m >= 0 {
		if matrix[m][n] == target {
			return true
		} else if matrix[m][n] < target {
			n++
		} else {
			m--
		}
	}
	return false
}
