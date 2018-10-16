package main

import (
	"fmt"
	"strings"
)

func main() {
	A := "AAAAA"
	fmt.Println(toLowerCase(A))
}
func toLowerCase(str string) string {

	return strings.ToLower(str)
}
