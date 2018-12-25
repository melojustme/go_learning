package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "race a car"
	b := isPalindrome(a)
	fmt.Println(b)

}
func isPalindrome(s string) bool {
	// fmt.Println(s)
	var a []byte
	var s1 = strings.ToLower(s)
	// fmt.Println(s1)
	for i := 0; i < len(s1); i++ {
		if 'a' <= s1[i] && s1[i] <= 'z' || '0' <= s1[i] && s1[i] <= '9' {
			a = append(a, s1[i])
		}
	}
	var l = len(a)
	for i := 0; i < l/2; i++ {
		if a[i] != a[l-i-1] {
			return false
		}
	}
	return true
}
