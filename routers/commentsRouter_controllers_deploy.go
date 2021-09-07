package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Mustang/controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["Mustang/controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/add",
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Mustang/controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["Mustang/controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Mustang/controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["Mustang/controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "Exec",
            Router: "/exec",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Mustang/controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["Mustang/controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
