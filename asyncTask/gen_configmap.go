package asyncTask

import (
	"Mustang/models"

	"github.com/RichardKnop/machinery/v1/tasks"
)

func GenConfigMap(id int64) error {
	_, err := models.DeployModel.GetById(id)
	if err != nil {
		return err
	}
	return nil
}

func GenConfigMapSignature(deploy *models.Deploy) *tasks.Signature {
	return &tasks.Signature{
		Name: "gen_configmap",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: deploy.Id,
			},
		},
	}
}
