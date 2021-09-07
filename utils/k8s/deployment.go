package k8s

import (
	"Mustang/models"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDeployment(deploy *models.Deploy) *appv1.Deployment {
	podLabels := map[string]string{
		"app": deploy.ProjectName,
	}
	deployment := &appv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploy.ProjectName,
			Namespace: deploy.EnvClusterBinding.Namespace,
		},
		Spec: appv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: podLabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: podLabels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            deploy.ProjectName,
							Image:           deploy.Image,
							ImagePullPolicy: "Always",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: int32(deploy.Port),
								},
							},
						},
					},
				},
			},
		},
	}
	return deployment
}
