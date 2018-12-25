package main

import (
	"fmt"
)

func main() {
	a := "loveleetcode"
	b := firstUniqChar(a)
	fmt.Println(b)

}

func firstUniqChar(s string) int {
	var h = map[byte]int{}
	for i := 0; i < len(s); i++ {
		h[s[i]] += 1
	}

	for i := 0; i < len(s); i++ {
		for k, v := range h {
			if v == 1 {
				if s[i] == k {
					return i
				}
			}
		}
	}

	return -1
}
