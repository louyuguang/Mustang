package machinery

import (
	"Mustang/app"
	"Mustang/asyncTask"
	"Mustang/models"
	"context"

	"github.com/RichardKnop/machinery/v1/tasks"
)

func TaskDeploy(deploy *models.Deploy) error {
	server := app.Machinery
	dockerBuildSig := asyncTask.DockerBuildSignature(deploy)
	genConfigmapSig := asyncTask.GenConfigMapSignature(deploy)
	deploySig := asyncTask.DeploySignature(deploy)

	chain, _ := tasks.NewChain(dockerBuildSig, genConfigmapSig, deploySig)
	_, err := server.SendChainWithContext(context.Background(), chain)
	if err != nil {
		return err
	}
	return nil
}
