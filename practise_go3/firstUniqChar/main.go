package main

import (
	"fmt"
)

func main() {
	var A = "leetcode"
	fmt.Println(firstUniqChar(A))
}
func firstUniqChar(s string) int {
	var h = make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		h[s[i]] += 1
	}
	// fmt.Println(h)
	for _, v := range h { //查询是否有不重复字符
		if v == 1 {
			// fmt.Println(k)
			break
		} else if v > 1 {
			continue
		} else {
			return -1
		}
	}
	//
	for i := 0; i < len(s); i++ {
		for k, v := range h {
			if v == 1 {
				if k == s[i] {
					return i
				}
			}
		}
	}
	return -1
}
