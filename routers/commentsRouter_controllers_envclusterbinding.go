package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Mustang/controllers/envclusterbinding:EnvClusterBindingController"] = append(beego.GlobalControllerRouter["Mustang/controllers/envclusterbinding:EnvClusterBindingController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/add",
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
