package main

import (
	"fmt"
)

func main() {
	a := []string{"a", "a", "l"}
	a = append(a[:0], a[1:]...)
	fmt.Println(a)
}
