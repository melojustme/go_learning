package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	id := ""
	a := WhichAgeByIdcard(id)
	fmt.Println(a)
}

//which age
func WhichAgeByIdcard(idcard string) string {
	fmt.Println(idcard)
	cyear, _ := strconv.Atoi(time.Now().Format("2006"))
	byear, _ := strconv.Atoi(idcard[6:10])
	return strconv.Itoa(cyear - byear)
}
