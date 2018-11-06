package main

import (
	"fmt"
	"strings"
)

func main() {
	// haystack := "aaaaa"
	// needle := "bba"
	haystack := "hello"
	needle := "ll"
	fmt.Println(haystack, needle)
	i := strStr(haystack, needle)
	fmt.Println(i)
}
func strStr(haystack string, needle string) int {
	if strings.Contains(haystack, needle) {
		a1 := strings.Index(haystack, needle)
		// fmt.Println(a1)
		return a1
	}
	return -1
}
