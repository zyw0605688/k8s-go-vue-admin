package api

import (
	"context"
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/config"
	"github.com/gin-gonic/gin"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceInfo struct {
	ServiceName             string `json:"service_name"`
	ServiceNamespace        string `json:"service_namespace"`
	ServicePort             int32  `json:"service_port"`
	ServiceNodePort         int32  `json:"service_node_port"`
	ServiceSelectorPodValue string `json:"service_selector_pod_value"`
}

func CreateService(c *gin.Context) {
	var info ServiceInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return
	}
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: info.ServiceName,
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{{
				Name:     info.ServiceName,
				Port:     info.ServicePort,
				NodePort: info.ServiceNodePort,
			},
			},
			Selector: map[string]string{
				"app": info.ServiceSelectorPodValue,
			},
			Type: apiv1.ServiceTypeNodePort,
		},
	}
	result, _ := config.ClientSet.CoreV1().Services(info.ServiceNamespace).Create(context.TODO(), service, metav1.CreateOptions{})
	c.JSON(200, gin.H{
		"message": result,
	})
}

func DeleteService(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")
	_ = config.ClientSet.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	c.JSON(200, gin.H{
		"message": name,
	})
}

func GetServiceList(c *gin.Context) {
	namespace := c.Query("namespace")
	result, _ := config.ClientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	c.JSON(200, gin.H{
		"message": result,
	})
}
