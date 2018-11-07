package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "dfuyalc"
	typed := "fuuyallc"

	// fmt.Println(name, typed)
	a := isLongPressedName(name, typed)
	fmt.Println(a)

}

func isLongPressedName(name string, typed string) bool {
	a := strings.Split(name, "")
	b := strings.Split(typed, "")
	var re bool
	l := len(b)
	for i := 0; i < len(a); i++ {
		s := a[i]
		// fmt.Println("a的数组", s)
		// fmt.Println("b的数组", b)

		for {
			if l > 0 { //如果长度为0,跳出
				if b[0] == s {
					// b = append(b[:l], b[l+1:]...)
					b = append(b[:0], b[1:]...)
					l -= 1
					break
				} else if b[0] == a[i-1] {
					// b = append(b[:l], b[l+1:]...)
					b = append(b[:0], b[1:]...)
					// fmt.Println("没有获取到值")
					l -= 1
					continue
				} else {
					return false
				}
			} else {
				// fmt.Println("mei有")
				// re = false
				return false
				// break
			}
		}
		re = true
	}

	return re
}
