package main

import (
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/config"
	"gitee.com/zyw0605688_admin/k8s-go-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitK8s()
	r := gin.Default()
	router.Setup(r)
	_ = r.Run(":7000")
}
