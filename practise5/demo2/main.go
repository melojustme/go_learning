package main

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

func main() {
	difficulty := 3
	prefix := strings.Repeat("0", difficulty)
	beego.Debug(prefix)

	nonce := fmt.Sprintf("%x", 9)

	beego.Debug(nonce)
}
