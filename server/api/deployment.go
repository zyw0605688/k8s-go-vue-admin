package api

import (
	"context"
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/config"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateDeployment(c *gin.Context) {
	// 构建参数
	deployment := &appsv1.Deployment{}
	namespace := "zyw"
	// 操作
	result, err := config.ClientSet.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func DeleteDeployment(c *gin.Context) {
	// 构建参数
	deployment := &appsv1.Deployment{}
	namespace := "zyw"
	// 操作
	err := config.ClientSet.AppsV1().Deployments(namespace).Delete(context.TODO(), deployment.Name, metav1.DeleteOptions{})
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "",
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
