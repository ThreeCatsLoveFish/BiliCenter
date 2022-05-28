package api

import (
	"net/http"
	"subcenter/application/awpush"

	"github.com/gin-gonic/gin"
)

// UpdateRelation update bilibili account relation
func UpdateRelation(c *gin.Context) {
	c.JSON(http.StatusOK, awpush.UpdateRelation())
}
