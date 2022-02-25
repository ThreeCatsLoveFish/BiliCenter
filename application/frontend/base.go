package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadFrontend(router *gin.Engine) {
	// load html template
	router.LoadHTMLGlob("template/*")

	// add router for frontend service
	router.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/push.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "push.html", nil)
	})
}
