package main

import (
	"fmt"

	"github.com/emirpasic/gods/utils"
)

func main() {
	var a = []interface{}{333, 22, 11, 312, 2}
	fmt.Println(a)
	utils.Sort(a, utils.IntComparator)
	fmt.Println(a)

}
