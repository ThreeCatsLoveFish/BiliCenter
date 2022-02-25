package api

import (
	"net/http"
	"subcenter/domain/push"

	"github.com/gin-gonic/gin"
)

// AddEndpoint update endpoint data
func AddEndpoint(c *gin.Context) {
	var endpoint push.Endpoint
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  5,
			"error": err,
		})
	}
	push.SetEndpoint(endpoint)
	c.JSON(http.StatusOK, endpoint)
}

// ListEndpoint list all endpoint data
func ListEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, push.GetEndpoint())
}
