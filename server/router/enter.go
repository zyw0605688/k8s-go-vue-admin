package router

import (
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/api"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	{
		r.GET("/node", api.GetNodeList)
	}
	{
		r.POST("/namespace", api.AddNameSpace)
		r.DELETE("/namespace", api.DeleteNameSpace)
		r.GET("/namespace", api.GetNameSpaceList)
	}

	{
		r.POST("/deployment", api.CreateDeployment)
		r.DELETE("/deployment", api.DeleteDeployment)
		r.PUT("/deployment", api.UpdateDeployment)
		r.GET("/deployment", api.GetDeploymentList)
	}

	{
		r.POST("/pod", api.GetPodList)
		r.DELETE("/pod", api.GetPodList)
		r.PUT("/pod", api.GetPodList)
		r.GET("/pod", api.GetPodList)
	}

	{
		r.POST("/service", api.GetServiceList)
		r.DELETE("/service", api.GetServiceList)
		r.PUT("/service", api.GetServiceList)
		r.GET("/service", api.GetServiceList)
	}

	{
		r.GET("/ingress", api.GetIngressList)
	}
}
