package api

import (
	"github.com/gin-gonic/gin"
)

// LoadApi add router for api service
func LoadApi(router *gin.Engine) {
	router.GET("/api/push/endpoint/list/", ListEndpoint)
	router.POST("/api/push/endpoint/add/", AddEndpoint)
}
