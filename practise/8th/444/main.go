package main

import (
	"time"

	"github.com/astaxie/beego"
)

func main() {
	etime := time.Now().Format("200601")
	beego.Debug(etime)
}
