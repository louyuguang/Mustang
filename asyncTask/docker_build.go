package asyncTask

import (
	"Mustang/models"

	"github.com/RichardKnop/machinery/v1/tasks"
)

func DockerBuild(id int64) error {
	deploy, err := models.DeployModel.GetById(id)
	if err != nil {
		return err
	}
	deploy.Image = "nginx:latest"
	if err = models.DeployModel.UpdateById(deploy); err != nil {
		return err
	}
	return nil
}

func DockerBuildSignature(deploy *models.Deploy) *tasks.Signature {
	return &tasks.Signature{
		Name: "docker_build",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: deploy.Id,
			},
		},
	}
}
