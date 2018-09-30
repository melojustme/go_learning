package main

import (
	"fmt"
)

func main() {
	var i = 3
	func(a int) {
		fmt.Println(a)
	}(i) //要传得值
}
