package api

import (
	"net/http"
	"subcenter/domain/push"

	"github.com/gin-gonic/gin"
)

// UpdateEndpoint update endpoint data
func UpdateEndpoint(c *gin.Context) {
	var endpoint push.Endpoint
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 4,
			"msg":  err,
		})
	}
	push.SetEndpoint(endpoint)
	c.JSON(http.StatusOK, endpoint)
}

// ListEndpoint list all endpoint data
func ListEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, push.GetEndpoint())
}
