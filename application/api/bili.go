package api

import (
	"net/http"
	"subcenter/application/awpush"

	"github.com/gin-gonic/gin"
)

// UpdateRelation update bilibili account relation
func UpdateRelation(c *gin.Context) {
	awpush.UpdateRelation()
	c.Data(http.StatusOK, "text/plain", []byte("Success!"))
}
