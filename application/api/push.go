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
		c.JSON(http.StatusOK, Resp{
			Code: Success,
			Msg:  err.Error(),
		})
		return
	}
	push.SetEndpoint(endpoint)
	c.JSON(http.StatusOK, Resp{
		Code: Success,
	})
}

// ListEndpoint list all endpoint data
func ListEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, Resp{
		Code: Success,
		Data: push.GetEndpoint(),
	})
}
