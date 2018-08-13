package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	a := RandomMa()
	fmt.Println(a)
}
func RandomMa() (ra string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := r.Intn(5)
	for {
		if a == 0 {
			a = r.Intn(5)
		} else {
			break
		}
	}
	ra1 := strconv.Itoa(a)
	ra2 := strconv.Itoa(r.Intn(100))
	return ra1 + "." + ra2
}
