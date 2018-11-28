package routers

import (
	"go_learning/practise_go3/demo2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/down1", "./down1")
	beego.SetStaticPath("/down2", "./down2")

	beego.Router("/", &controllers.MainController{})
	beego.Router("/list", &controllers.ListController{})

}
