package main

import (
	"flag"
	"fmt"
)

var workers = flag.Int("r", 1, "concurrent processing ,默认 1 .")

var workers2 = flag.String("st", "哈哈哈哈", "concurrent processing ,default 哈哈哈哈 .")

func main() {
	flag.Parse()
	fmt.Println(*workers)
	fmt.Println(*workers2)

}
