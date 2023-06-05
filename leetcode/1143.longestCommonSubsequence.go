package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	res := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		res[i] = make([]int, len(text2)+1)
	}
	var max func(a, b int) int
	max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				res[i+1][j+1] = res[i][j] + 1
			} else {
				res[i+1][j+1] = max(res[i+1][j], res[i][j+1])
			}
		}
	}
	return res[len(text1)][len(text2)]
}
