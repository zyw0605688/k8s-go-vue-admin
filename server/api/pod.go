package api

import (
	"context"
	"fmt"
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/config"
	"github.com/gin-gonic/gin"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func GetPodList(c *gin.Context) {
	namespace := c.Query("namespace")
	result := &v1.PodList{}
	err := config.RestClient.Get().
		Namespace(namespace).
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"message": result,
	})
}
