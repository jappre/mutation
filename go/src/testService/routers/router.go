package routers

import (
	"github.com/astaxie/beego"
	"testService/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Get;post:ModifyEmail")
	// beego.Router("/user/modify_email", &controllers.ModifyEmailController{})
}
