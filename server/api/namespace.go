package api

import (
	"context"
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/config"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ReqNamespace struct {
	Namespace string `json:"namespace"`
}

func AddNameSpace(c *gin.Context) {
	var req ReqNamespace
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	name := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Namespace,
		},
	}

	result, _ := config.ClientSet.CoreV1().Namespaces().Create(context.TODO(), name, metav1.CreateOptions{})
	c.JSON(200, gin.H{
		"message": result,
	})
}

func DeleteNameSpace(c *gin.Context) {
	result, _ := config.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	c.JSON(200, gin.H{
		"message": result,
	})
}

func GetNameSpaceList(c *gin.Context) {
	result, _ := config.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	c.JSON(200, gin.H{
		"message": result,
	})
}
