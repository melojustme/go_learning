package main

import (
	"fmt"
	"strings"
)

func main() {
	// a := 38
	J := "aA"
	S := "aAAbbbb"
	b := numJewelsInStones(J, S)

	fmt.Println(b)
}
func numJewelsInStones(J string, S string) int {
	var i int
	arr := strings.Split(J, "")
	for _, v := range arr {
		a := strings.Count(S, v)
		i += a
	}
	return i
}
