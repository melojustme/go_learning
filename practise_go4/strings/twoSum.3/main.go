package main

import (
	"fmt"
)

func main() {
	a := "anagram"
	t := "nagaram"

	b := isAnagram(a, t)
	fmt.Println(b)

}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var h, a = map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s); i++ {
		h[s[i]] += 1
	}
	for i := 0; i < len(t); i++ {
		a[t[i]] += 1
	}
	fmt.Println(h, a)
	if len(h) != len(a) {
		return false
	}
	for k, _ := range h {
		if h[k] != a[k] {
			return false
		}
	}
	return true
}
