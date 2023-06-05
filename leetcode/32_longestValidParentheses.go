package leetcode

/*
最长有效括号
题目：https://leetcode.cn/problems/longest-valid-parentheses/
题解：
*/

// left right
func longestValidParentheses1(s string) int {
	left, right := 0, 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxlen := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxlen = max(maxlen, 2*left)
		} else if right > left {
			right, left = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxlen = max(maxlen, left*2)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxlen
}

// 利用栈
func longestValidParentheses(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxLen := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}
