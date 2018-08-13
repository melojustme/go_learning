package main

import (
	"time"

	"github.com/astaxie/beego"
)

func main() {
	birthDay18 := time.Now().AddDate(-18, 0, 0).Format("2006-01-02")
	birthDay19 := time.Now().AddDate(-20, 0, 0).Format("2006-01-02")
	beego.Debug(birthDay18, birthDay19)
	birthDay20 := time.Now().AddDate(-20, 0, 0).Format("2006-01-02")
	birthDay29 := time.Now().AddDate(-30, 0, 0).Format("2006-01-02")
	beego.Debug(birthDay20, birthDay29)
	birthDay30 := time.Now().AddDate(-30, 0, 0).Format("2006-01-02")
	birthDay39 := time.Now().AddDate(-40, 0, 0).Format("2006-01-02")
	beego.Debug(birthDay30, birthDay39)
	birthDay40 := time.Now().AddDate(-30, 0, 0).Format("2006-01-02")
	birthDay49 := time.Now().AddDate(-40, 0, 0).Format("2006-01-02")
	beego.Debug(birthDay40, birthDay49)
	birthDay50 := time.Now().AddDate(-50, 0, 0).Format("2006-01-02")
	birthDay59 := time.Now().AddDate(-60, 0, 0).Format("2006-01-02")
	beego.Debug(birthDay50, birthDay59)
	birthDay60 := time.Now().AddDate(-60, 0, 0).Format("2006-01-02")
	birthDay69 := time.Now().AddDate(-70, 0, 0).Format("2006-01-02")
	beego.Debug(birthDay60, birthDay69)

}
