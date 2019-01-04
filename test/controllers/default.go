package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.html"
}
func (c *MainController) Post() {
	defer c.ServeJSON()
	id := c.GetString("id")
	data := c.GetString("data")
	beego.Debug(id)
	beego.Debug(data)
	var a Data
	json.Unmarshal([]byte(data), &a)
	fmt.Println(a)
	c.Data["json"] = 200
}

type Data struct {
	Name string `json:"name"`
	Text string `json:"text"`
}
