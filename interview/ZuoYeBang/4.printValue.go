package ZuoYeBang

import "strings"

/*
from 作业帮
输入： a=1&b=2&b=3&b=4&c=5&c=6&a=22
输出：a:{1, 22}
	 b:{2, 3, 4}
	 c:{5, 6}
*/

func GetValue(str string) map[string][]string {
	mp := make(map[string][]string, 0)
	if len(str) == 0 {
		return mp
	}
	strArr := strings.Split(str, "&")
	for i := 0; i < len(strArr); i++ {
		s := strings.Split(strArr[i], "=")
		if _, ok := mp[s[0]]; !ok {
			mp[s[0]] = []string{s[1]}
		} else {
			tempRes := mp[s[0]]
			tempRes = append(tempRes, s[1])
			mp[s[0]] = tempRes
		}
	}
	return mp
}
