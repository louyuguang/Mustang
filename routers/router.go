package routers

import (
	"Mustang/controllers/auth"
	"Mustang/controllers/base"
	"Mustang/controllers/user"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&base.BaseController{})
	beego.Include(&auth.AuthController{})
	nsWithUser := beego.NewNamespace("/user",
		beego.NSInclude(&user.UserController{}),
	)
	beego.AddNamespace(nsWithUser)
}
