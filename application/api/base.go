package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoadApi add router for api service
func LoadApi(router *gin.Engine) {
	// DEMO for api data
	router.GET("/api/demo/get/:data", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"data": c.Param("data"),
		})
	})
	// TODO: edit to json for set, input json, add
	router.POST("/api/push/endpoint/add/")
	// TODO: edit to json for get, input string, obtain
	router.GET("/api/push/endpoint/list/:name")
}
