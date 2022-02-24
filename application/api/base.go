package api

import (
	"github.com/gin-gonic/gin"
)

func LoadApi(router *gin.Engine) {
	// add router for api service
	// TODO: edit to json for set
	router.POST("/api/push/endpoint/add/")
	// TODO: edit to json for get
	router.GET("/api/push/endpoint/list/")
}
