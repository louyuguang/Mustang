package machinery

import (
	"Mustang/app"
	"Mustang/asyncTask"
	"Mustang/utils/logs"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

func CreateMachinery() *machinery.Server {
	if app.Machinery != nil {
		return app.Machinery
	}
	cnf, err := config.NewFromYaml("conf/machinery.yaml", false)
	if err != nil {
		logs.Critical("read config failed", err)
	}

	MachineryServer, err := machinery.NewServer(cnf)
	if err != nil {
		logs.Critical("start server failed", err)
	}
	tasksList := map[string]interface{}{
		"docker_build":  asyncTask.DockerBuild,
		"gen_configmap": asyncTask.GenConfigMap,
		"deploy":        asyncTask.Deploy,
		//"deploy": asyncTask.
	}
	MachineryServer.RegisterTasks(tasksList)
	app.Machinery = MachineryServer
	return MachineryServer
}
