package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subcenter/domain/push"
)

// LoadApi add router for api service
func LoadApi(router *gin.Engine) {
	// DEMO for api data
	router.GET("/api/demo/get/:data", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"data": c.Param("data"),
		})
	})
	// edit to json for set, input json, add
	router.POST("/api/push/endpoint/add/", SetHandler)
	//  edit to json for get, input string, obtain
	router.GET("/api/push/endpoint/list/", GetHandler)
}

func SetHandler(c *gin.Context) {
	var endpnt push.Endpoint
	if err := c.ShouldBindJSON(&endpnt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	push.SetEndpoint(endpnt)
	c.JSON(http.StatusOK, endpnt)

}

func GetHandler(c *gin.Context) {
	info := push.GetEndpoint()
	c.JSON(http.StatusOK, info)
}
