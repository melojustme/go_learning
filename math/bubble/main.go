package main

import (
	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	var a = []int{4, 5, 6, 3, 2, 1}
	fmt.Println(a)
	b := bulle(a)
	fmt.Println(b)
}

//冒泡
func bulle(a []int) []int {
	for i := len(a) - 1; i > 0; i-- {
		beego.Debug("a", i)
		for j := 0; j < i; j++ {
			beego.Debug("排序中b", j)
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
			fmt.Println("结果", a)
		}
	}
	return a
}
