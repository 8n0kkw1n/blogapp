package routers

import (
	"hwgo/5/blog_app_orm/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
