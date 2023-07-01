package api

import (
	"context"
	"fmt"
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/config"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/networking/v1"
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

	// 构建参数开始
	ingress := &v1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: info.IngressName,
		},
		Spec: v1.IngressSpec{
			IngressClassName: Traefik,
			Rules: []v1.IngressRule{
				{
					Host: info.IngressHost,
					IngressRuleValue: v1.IngressRuleValue{
						HTTP: &v1.HTTPIngressRuleValue{
							Paths: []v1.HTTPIngressPath{
								{
									Path: "/",
									Backend: v1.IngressBackend{
										Service: &v1.IngressServiceBackend{
											Name: info.IngressServiceName,
											Port: v1.ServiceBackendPort{
												Number: info.IngressServicePort,
											},
										},
									},
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

	result, _ := config.ClientSet.NetworkingV1().Ingresses(info.IngressNamespace).Create(context.TODO(), ingress, metav1.CreateOptions{})
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
