package asyncTask

import (
	"Mustang/models"
	"Mustang/utils/k8s"
	"Mustang/utils/logs"
	"context"
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/tasks"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Deploy(id int64) error {
	deploy, err := models.DeployModel.GetById(id)
	if err != nil {
		return err
	}
	clusterIds := strings.Split(deploy.EnvClusterBinding.ClusterIds, ",")
	for _, v := range clusterIds {
		clusterId, _ := strconv.Atoi(v)
		cluster, err := models.ClusterModel.GetById(int64(clusterId))
		if err != nil {
			return err
		}
		clientSet, err := k8s.GetClientSetFromByte([]byte(cluster.KubeConfig))
		if err != nil {
			return err
		}
		deployment := k8s.GetDeployment(deploy)
		_, err = clientSet.AppsV1().Deployments(deploy.EnvClusterBinding.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		logs.Info("Env：%s ClusterId: %d deploy %s success.", deploy.EnvClusterBinding.EnvName, int64(clusterId), deploy.ProjectName)
	}
	return nil
}

func DeploySignature(deploy *models.Deploy) *tasks.Signature {
	return &tasks.Signature{
		Name: "deploy",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: deploy.Id,
			},
		},
	}
}
