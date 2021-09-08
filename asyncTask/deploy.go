package asyncTask

import (
	"Mustang/models"
	"Mustang/utils/k8s"
	"Mustang/utils/logs"
	"context"
	"strconv"
	"strings"

	"github.com/RichardKnop/machinery/v1/tasks"
	"k8s.io/apimachinery/pkg/api/errors"
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
		// get k8s clientSet and deploymentClient
		clientSet, err := k8s.GetClientSetFromByte([]byte(cluster.KubeConfig))
		deploymentClient := clientSet.AppsV1().Deployments(deploy.EnvClusterBinding.Namespace)
		if err != nil {
			return err
		}
		// generate deployment object
		deployment := k8s.GetDeployment(deploy)
		// get project's deployment object from k8s cluster
		// if deployment exist, exec update(), if not exec create()
		if oldDeployment, err := deploymentClient.Get(context.TODO(), deployment.Name, metav1.GetOptions{}); err != nil {
			if errors.IsNotFound(err) {
				_, err = deploymentClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
			} else {
				return err
			}
		} else {
			deployment.Spec.Replicas = oldDeployment.Spec.Replicas
			if _, err := deploymentClient.Update(context.TODO(), deployment, metav1.UpdateOptions{}); err != nil {
				return err
			}
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
