package api

import (
	"context"
	"fmt"
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/config"
	"github.com/gin-gonic/gin"
	newworkV1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IngressInfo struct {
	IngressName        string `json:"ingress_name"`
	IngressNamespace   string `json:"ingress_namespace"`
	IngressHost        string `json:"ingress_host"`
	IngressServiceName string `json:"ingress_service_name"`
	IngressServicePort int32  `json:"ingress_service_port"`
}

func CreateIngress(c *gin.Context) {
	var info IngressInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return
	}

	traefik := "traefik"
	var Traefik *string = &traefik
	var pathType = newworkV1.PathTypePrefix

	// 构建参数开始
	ingress := &newworkV1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: info.IngressName,
		},
		Spec: newworkV1.IngressSpec{
			IngressClassName: Traefik,
			Rules: []newworkV1.IngressRule{
				{
					Host: info.IngressHost,
					IngressRuleValue: newworkV1.IngressRuleValue{
						HTTP: &newworkV1.HTTPIngressRuleValue{
							Paths: []newworkV1.HTTPIngressPath{
								{
									Path: "/",
									Backend: newworkV1.IngressBackend{
										Service: &newworkV1.IngressServiceBackend{
											Name: info.IngressServiceName,
											Port: newworkV1.ServiceBackendPort{
												Number: info.IngressServicePort,
											},
										},
									},
									PathType: &pathType,
								},
							},
						},
					},
				},
			},
		},
	}
	// 构建参数结束
	fmt.Println(">>>", ingress)

	result, err := config.ClientSet.NetworkingV1().Ingresses(info.IngressNamespace).Create(context.TODO(), ingress, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("错误信息：", err)
	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func DeleteIngress(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")
	_ = config.ClientSet.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	c.JSON(200, gin.H{
		"message": name,
	})
}

func GetIngressList(c *gin.Context) {
	namespace := c.Query("namespace")
	result, _ := config.ClientSet.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	c.JSON(200, gin.H{
		"message": result,
	})
}
