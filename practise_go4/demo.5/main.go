package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []string{"8", "2", "3", "4"}
	fmt.Println(GetMaxNum(a))
}
func GetMaxNum(a []string) string {
	var b, c int
	for _, v := range a {
		fmt.Println(v)
		b, _ = strconv.Atoi(v)
		if b > c {
			c = b
		}

	}
	return strconv.Itoa(c)
}
