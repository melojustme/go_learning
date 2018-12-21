package main

import (
	"fmt"
	"strings"
)

func main() {
	var A = "aa"
	var B = "a"
	fmt.Println(rotateString(A, B))
}

func rotateString(A string, B string) bool {
	fmt.Println(A, B)
	if len(A) != len(B) {
		return false
	}
	if strings.Contains(B+B, A) {
		return true
	}
	return false
}
