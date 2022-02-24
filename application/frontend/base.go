package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewFrontend() *gin.Engine {
	router := gin.Default()

	// load html template
	router.LoadHTMLGlob("template/*")

	// add router for frontend service
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	return router
}
