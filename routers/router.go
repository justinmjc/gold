package routers

import (
	"gold/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	api := beego.NewNamespace("/v1",
		//user login
		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.MainController{}, "get:Get"),
		),
		beego.NSNamespace("/oss",
			beego.NSRouter("/", &controllers.MainController{}, "get:Get"),
		),
	)

	beego.AddNamespace(api)
}
