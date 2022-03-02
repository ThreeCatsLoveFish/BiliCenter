package api

import (
	"net/http"
	"subcenter/application/awpush"

	"github.com/gin-gonic/gin"
)

// UpdateBili update bili account
func UpdateBili(c *gin.Context) {
	if result := awpush.UpdateRelation(); result == nil {
		c.JSON(http.StatusOK, Resp{
			Code: Success,
		})
	} else {
		c.JSON(http.StatusOK, Resp{
			Code: ErrInternal,
			Data: result,
		})
	}
}
