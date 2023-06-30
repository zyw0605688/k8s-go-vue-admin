package api

import (
	"context"
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/config"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetServiceList(c *gin.Context) {
	namespace := c.Query("namespace")
	result, _ := config.ClientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	c.JSON(200, gin.H{
		"message": result,
	})
}
