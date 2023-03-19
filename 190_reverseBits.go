package leetcode

/*
题目：颠倒的二进制
颠倒给定的 32 位无符号整数的二进制位。
题解：https://leetcode.cn/problems/reverse-bits/solution/by-br41n10-295v/
*/
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		bit := num & 1
		num >>= 1
		res <<= 1
		if bit == 1 {
			res++
		}
	}
	return res
}
