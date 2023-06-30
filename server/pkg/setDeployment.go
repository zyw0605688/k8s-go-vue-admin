package pkg

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentInfo struct {
	PodNamespace  string  `json:"pod_namespace"`
	PodName       string  `json:"pod_name"`
	PodCpuMax     float32 `json:"pod_cpu_max"`
	PodReplicas   int32   `json:"pod_replicas"`
	PodMemoryMax  float32 `json:"pod_memory_max"`
	PodPullPolicy string  `json:"pod_pull_policy"`
	PodRestart    string  `json:"pod_restart"`
	PodType       string  `json:"pod_type"`
	PodImage      string  `json:"pod_image"`
}

func SetDeployment() {
	var info DeploymentInfo
	deployment := &appsv1.Deployment{}
	deployment.TypeMeta = metav1.TypeMeta{
		Kind:       "deployment",
		APIVersion: "v1",
	}
	deployment.ObjectMeta = metav1.ObjectMeta{
		Name:      info.PodName,
		Namespace: info.PodNamespace,
		Labels: map[string]string{
			"app":    info.PodName,
			"author": "zyw",
		},
	}
	deployment.Name = info.PodName
	deployment.Spec = appsv1.DeploymentSpec{
		Replicas: &info.PodReplicas,
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app-name": info.PodName,
			},
			MatchExpressions: nil,
		},
		Template: v1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app-name": info.PodName,
				},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  info.PodName,
						Image: info.PodImage,
					},
				},
			},
		},
		Strategy:                appsv1.DeploymentStrategy{},
		MinReadySeconds:         0,
		RevisionHistoryLimit:    nil,
		Paused:                  false,
		ProgressDeadlineSeconds: nil,
	}
}
