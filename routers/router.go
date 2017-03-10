package routers

import (
	"hblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/article/add", &controllers.RestController{}, "*:Add")
	beego.Router("/api/article/list", &controllers.RestController{}, "*:List")
}
