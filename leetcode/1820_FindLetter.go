package leetcode

/*
from LintCode
给定一个字符串str，返回字符串中字母顺序最大的而且同时在字符串中出现大写和小写的字母。
如果不存在这样的字母，返回‘~‘。
输入:"aAbBcD"
输出:'B'
解释：因为c和D没有大小写同时出现，A和B都有大小写，但是B比A大，所以返回B。
*/

func FindLetter(str string) byte {
	// Write your code here.
	if len(str) == 0 {
		return '~'
	}
	mpA := make(map[byte]int, 26)
	mpa := make(map[byte]int, 26)
	for i := 0; i < len(str); i++ {
		ch := str[i]
		if ch >= 'a' && ch <= 'z' {
			mpa[ch]++
		}
		if ch >= 'A' && ch <= 'Z' {
			mpA[ch]++
		}
	}
	for i := 'Z'; i >= 'A'; i-- {
		_, okA := mpA[byte(i)]
		_, oka := mpa[byte(i)+32]
		if oka && okA {
			return byte(i)
		}
	}
	return '~'
}
