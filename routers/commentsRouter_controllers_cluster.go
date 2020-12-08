package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Mustang/controllers/cluster:ClusterController"] = append(beego.GlobalControllerRouter["Mustang/controllers/cluster:ClusterController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/add",
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Mustang/controllers/cluster:ClusterController"] = append(beego.GlobalControllerRouter["Mustang/controllers/cluster:ClusterController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Mustang/controllers/cluster:ClusterController"] = append(beego.GlobalControllerRouter["Mustang/controllers/cluster:ClusterController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Mustang/controllers/cluster:ClusterController"] = append(beego.GlobalControllerRouter["Mustang/controllers/cluster:ClusterController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/update/:id",
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
