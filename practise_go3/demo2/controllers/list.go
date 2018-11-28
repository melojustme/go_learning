package controllers

import (
	"go_learning/practise_go3/demo2/models"
	"io/ioutil"

	"github.com/astaxie/beego"
)

type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {
	// var a []string

	path := "./down1"
	//遍历打印所有的文件名
	files, _ := ioutil.ReadDir(path)
	// var a []models.File
	a := make([]models.File, len(files))
	// for _, f := range files {
	// 	a.Name = append(a.Name, "http://localhost:8080/down1/"+f.Name())
	// }
	for i := 0; i < len(files); i++ {
		a[i].Name = files[i].Name()
		a[i].Src = "http://localhost:8080/down1/" + files[i].Name()

	}

	beego.Debug("地址", a)
	c.Data["list"] = a
	c.TplName = "video.html"
}
