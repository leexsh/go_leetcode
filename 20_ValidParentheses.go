package leetcode

// 有效的括号
// 切片模拟栈
// https://leetcode.cn/problems/valid-parentheses/

func isValid(s string) bool {
	if len(s)&1 == 0 {
		return false
	}
	var res []string
	for _, i := range s {
		temp_str := string(i)
		if i == '[' || i == '(' || i == '{' {
			res = append(res, temp_str)
		} else {
			if len(res) == 0 {
				return false
			}
			top := res[len(res)-1]
			if top == "(" && temp_str != ")" {
				return false
			} else if top == "{" && temp_str != "}" {
				return false
			} else if top == "[" && temp_str != "]" {
				return false
			}
			res = res[:len(res)-1]
		}
	}
	return len(res) == 0
}
