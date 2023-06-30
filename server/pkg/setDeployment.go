package pkg

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

type DeploymentInfo struct {
	PodName      string `json:"pod_name"`      // 名字
	PodNamespace string `json:"pod_namespace"` // 命名空间
	PodImage     string `json:"pod_image"`     // 镜像
	PodPort      int32  `json:"pod_port"`      // 端口号
	PodReplicas  int32  `json:"pod_replicas"`  // 副本数
}

func SetDeployment(info *DeploymentInfo) *appsv1.Deployment {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(info.PodReplicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": info.PodName,
				},
			},

			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": info.PodName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            info.PodName,
							Image:           info.PodImage,
							ImagePullPolicy: "IfNotPresent",
							Ports: []apiv1.ContainerPort{
								{
									Name:          info.PodName,
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: info.PodPort,
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
