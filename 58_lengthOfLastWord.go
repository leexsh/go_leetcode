package leetcode

/*
题目：最后一个单词的长度
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

题解：

*/

func lengthOfLastWord(s string) int {
	index := len(s) - 1

	for ; s[index] == ' '; index-- {

	}
	res := 0
	for ; index >= 0 && s[index] != ' '; index-- {
		res++
	}
	return res
}
