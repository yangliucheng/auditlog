package routers

import (
	"auditlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/auditlog/v1/writelog", &controllers.MainController{}, "POST:WriteLog")
	beego.Router("/auditlog/v1/getlog", &controllers.MainController{}, "GET:GetLog")
}
