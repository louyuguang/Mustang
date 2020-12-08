package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Mustang/controllers/user:UserController"] = append(beego.GlobalControllerRouter["Mustang/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/add",
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Mustang/controllers/user:UserController"] = append(beego.GlobalControllerRouter["Mustang/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Mustang/controllers/user:UserController"] = append(beego.GlobalControllerRouter["Mustang/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "Detail",
            Router: "/detail/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Mustang/controllers/user:UserController"] = append(beego.GlobalControllerRouter["Mustang/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Mustang/controllers/user:UserController"] = append(beego.GlobalControllerRouter["Mustang/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/update/:id",
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
