package main

import (
	"fmt"
)

func main() {
	cgPcode := []string{"1", "2", "1", "3", "1", "2", "1"}
	Comparison(cgPcode)
}

func Comparison(cp []string) {
	a := make(map[string]int, 0)
	var b []string
	for i := 0; i < len(cp); i++ {
		a[cp[i]] += 1
		if a[cp[i]] > 1 {
			b = append(b, cp[i])
		}
	}
	// fmt.Println(a)
	for k, v := range a {
		// fmt.Println(k, v)
		if v > 1 {
			fmt.Println(k)
		}
	}
}
