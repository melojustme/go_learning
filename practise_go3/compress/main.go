package main

import (
	"fmt"
	"strconv"
)

func main() {
	var A = "aaaaabbccc"
	fmt.Println(compress([]byte(A)))
}
func compress(chars []byte) int {
	h := make(map[byte]int, 0)
	var a []string
	for i := 0; i < len(chars); i++ {
		if i+1 < len(chars) {
			if chars[i] == chars[i+1] {
				h[chars[i]] += 1
			}
		}
	}
	fmt.Println(h)
	for k, v := range h {
		fmt.Println(k, v)
		a = append(a, string(k))
		a = append(a, strconv.Itoa(v))
	}
	fmt.Println(a)
	return 0
}
