package main

import "fmt"

func main() {
	a := "hello"
	fmt.Println(a)
	b := reverseString(a)
	fmt.Println(b)
}

func reverseString(s string) string {
	var a []byte
	for i := 0; i < len(s); i++ {
		// fmt.Println(s[i])
		a = append(a, s[i])
	}
	var b []byte
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
		// fmt.Println(a[i])
	}
	// fmt.Println(b)
	return string(b)
}
