package api

import (
	"net/http"
	"subcenter/domain/push"

	"github.com/gin-gonic/gin"
)

// LoadApi add router for api service
func LoadApi(router *gin.Engine) {
	// Add endpoint data
	router.POST("/api/push/endpoint/add/", SetHandler)
	// List endpoint data
	router.GET("/api/push/endpoint/list/", GetHandler)
}

func SetHandler(c *gin.Context) {
	var endpoint push.Endpoint
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	push.SetEndpoint(endpoint)
	c.JSON(http.StatusOK, endpoint)
}

func GetHandler(c *gin.Context) {
	info := push.GetEndpoint()
	c.JSON(http.StatusOK, info)
}
