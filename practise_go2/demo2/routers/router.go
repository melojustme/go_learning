package routers

import (
	"go_learning/practise_go2/demo2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/down1", "./down1")
	beego.Router("/", &controllers.MainController{})
}
