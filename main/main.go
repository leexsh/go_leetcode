package main

import (
	"fmt"
	"leexsh/leetcode/interview/ZuoYeBang"
)

func main() {
	str := "a=1&b=2&b=3&b=4&c=5&c=6&a=22"
	res := ZuoYeBang.GetValue(str)
	fmt.Println(res)
	for key, value := range res {
		fmt.Println(key, "-- ", value)
	}
}
