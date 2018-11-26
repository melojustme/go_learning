package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func main() {
	a := ULID()
	fmt.Println(a)
}
func ULID() string {
	t := time.Now()
	fmt.Println(t)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	// Output: 0000XSNJG0MQJHBF4QX1EFD6Y3
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
