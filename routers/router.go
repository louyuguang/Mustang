package routers

import (
	"Mustang/controllers/auth"
	"Mustang/controllers/base"
	"Mustang/controllers/cluster"
	"Mustang/controllers/deploy"
	"Mustang/controllers/envclusterbinding"
	"Mustang/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	//// Beego注解路由代码生成规则和程序运行路径相关，需要改写一下避免产生不一致的文件名
	//if beego.BConfig.RunMode == "dev" && path.Base(beego.AppPath) == "GoLand" {
	//	beego.AppPath = path.Join(path.Dir(beego.AppPath), "./")
	//}

	beego.Include(&base.BaseController{})
	beego.Include(&auth.AuthController{})
	nsWithUser := beego.NewNamespace("/user",
		beego.NSInclude(&user.UserController{}),
	)
	nsWithCluster := beego.NewNamespace("/cluster",
		beego.NSInclude(&cluster.ClusterController{}),
	)
	nsWithDeploy := beego.NewNamespace("/deploy",
		beego.NSInclude(&deploy.DeployController{}),
	)
	nsWithEnv := beego.NewNamespace("/env",
		beego.NSInclude(&envclusterbinding.EnvClusterBindingController{}),
	)
	beego.AddNamespace(nsWithUser)
	beego.AddNamespace(nsWithCluster)
	beego.AddNamespace(nsWithDeploy)
	beego.AddNamespace(nsWithEnv)
}
