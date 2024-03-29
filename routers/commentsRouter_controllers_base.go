package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Mustang/controllers/base:BaseController"] = append(beego.GlobalControllerRouter["Mustang/controllers/base:BaseController"],
        beego.ControllerComments{
            Method: "Index",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
