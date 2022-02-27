package api

import (
	"net/http"
	"subcenter/application/awpush"

	"github.com/gin-gonic/gin"
)

// UpdateBili update bili account
func UpdateBili(c *gin.Context) {
	c.JSON(http.StatusOK, awpush.UpdateRelation())
}
