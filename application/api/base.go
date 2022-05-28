package api

import (
	"github.com/gin-gonic/gin"
)

// LoadApi add router for api service
func LoadApi(router *gin.Engine) {
	// Bili account helper
	router.GET("/api/bili/account/clean/", ListHosts)

	// Hadoop Hosts config
	router.GET("/api/hadoop/hosts/list/", ListHosts)
	router.POST("/api/hadoop/hosts/update/", UpdateHosts)
	router.POST("/api/hadoop/hosts/reset/", ResetHosts)
}
