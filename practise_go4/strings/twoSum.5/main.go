package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "racar"
	b := "ar"
	c := strStr(a, b)
	fmt.Println(c)

}
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
