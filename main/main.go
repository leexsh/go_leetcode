package main

import (
	"fmt"
	"leexsh/leetcode/interview"
	"strings"
)

type ListNode struct {
	val  int
	next *ListNode
}
type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	pre := &ListNode{
		val:  -1,
		next: head,
	}
	cur, L, R := pre, pre, pre
	for cur.next != nil {
		L, R = cur.next, cur.next
		for R.next != nil && L.val == R.next.val {
			R = R.next
		}
		if L == R {
			cur = cur.next
		} else {
			cur.next = R.next
		}
	}
	return pre.next
}
func stringToFloat64(s string) (float64, error) {
	// 移除字符串中的空格
	s = strings.TrimSpace(s)

	// 如果字符串为空，则返回错误
	if s == "" {
		return 0, fmt.Errorf("empty string")
	}

	// 定义正负号标志
	negative := false

	// 如果字符串以负号开头，则将负号标志设置为 true，并去掉负号
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	var result float64
	var decimalSeen bool
	var decimalPlaces int

	// 遍历字符串中的字符
	for _, char := range s {
		// 如果是数字字符，则转换为数字并加入到结果中
		if char >= '0' && char <= '9' {
			result = result*10 + float64(char-'0')
			if decimalSeen {
				decimalPlaces++
			}
		} else if char == '.' {
			// 如果遇到小数点，则将小数点标志设置为 true
			if decimalSeen {
				return 0, fmt.Errorf("multiple decimal points")
			}
			decimalSeen = true
		} else {
			// 如果遇到非法字符，则返回错误
			return 0, fmt.Errorf("invalid character: %c", char)
		}
	}

	// 如果小数点标志为 true，则处理小数部分
	if decimalSeen {
		for i := 0; i < decimalPlaces; i++ {
			result /= 10
		}
	}

	// 如果是负数，则取负值
	if negative {
		result = -result
	}

	return result, nil
}
func main() {
	s1 := "123.023"
	s2 := ".1"
	s3 := "-32.23"
	s4 := "-123"
	strArr := []string{s1, s2, s3, s4}
	for i := 0; i < len(strArr); i++ {
		fmt.Println(interview.StringToFloat(strArr[i]))
		fmt.Println(stringToFloat64(strArr[i]))
	}
}
