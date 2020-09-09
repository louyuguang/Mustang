package routers

import (
	"Mustang/controllers"
	"Mustang/controllers/auth"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Include(&controllers.BaseController{})
	beego.Router("/", &controllers.BaseController{}, "get:Index")
	beego.Router("/login", &auth.AuthController{}, "get:Login;post:Login")
	beego.Router("/logout", &auth.AuthController{}, "get:Logout")
	ns := beego.NewNamespace("/user",
		beego.NSRouter("/add", &controllers.UserController{}, "get:Add;post:Add"),
		beego.NSRouter("/detail", &controllers.UserController{}, "get:Detail"),
		beego.NSRouter("/list", &controllers.UserController{}, "get:List"),
	)
	beego.AddNamespace(ns)
}
