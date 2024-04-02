package interview

import (
	"strings"
)

/*
string 字符串转为float64
"123.023" -> 123.023
".1" -> 0.1
"-32.23" -> -32.23
"-123" -> -123

from zoom
*/

func StringToFloat(input string) float64 {
	// 去掉空格
	strings.TrimSpace(input)
	flag := 1
	isAfterDot := false
	var res float64 = 0
	var res1 float64 = 0
	var bit float64 = 0.1
	for i := 0; i < len(input); i++ {
		if input[i] == '-' {
			flag = -1
		} else if input[i] == '.' {
			isAfterDot = true
		} else if isAfterDot {
			num := input[i] - '0'
			res1 = res1 + float64(num)*bit
			bit *= 0.1
		} else if !isAfterDot {
			num := input[i] - '0'
			res = res*10 + float64(num)
		}
	}
	return float64((res + res1) * float64(flag))
}

func stringToFloat642(s string) float64 {
	// 移除字符串中的空格
	s = strings.TrimSpace(s)

	// 如果字符串为空，则返回错误
	if s == "" {
		return 0
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
				return 0
			}
			decimalSeen = true
		} else {
			// 如果遇到非法字符，则返回错误
			return 0
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

	return result
}
