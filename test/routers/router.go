package routers

import (
	"go_learning/test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &controllers.MainController{})
}
