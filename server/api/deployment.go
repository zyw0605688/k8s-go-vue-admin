package api

import (
	"context"
	"fmt"
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/config"
	"github.com/gin-gonic/gin"
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

func CreateDeployment(c *gin.Context) {
	// 构建参数
	var info DeploymentInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: info.PodName,
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

	// 操作
	result, err := config.ClientSet.AppsV1().Deployments(info.PodNamespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func DeleteDeployment(c *gin.Context) {
	// 构建参数
	namespace := c.Query("namespace")
	name := c.Query("name")

	err := config.ClientSet.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": name,
	})
}

func UpdateDeployment(c *gin.Context) {
	// 构建参数
	deployment := &appsv1.Deployment{}
	namespace := "zyw"
	// 操作
	result, err := config.ClientSet.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func GetDeploymentList(c *gin.Context) {
	namespace := c.Query("namespace")
	result, _ := config.ClientSet.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	c.JSON(200, gin.H{
		"message": result,
	})
}
