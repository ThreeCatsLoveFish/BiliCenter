package api

import (
	"net/http"
    "subcenter/domain/push"
	"github.com/gin-gonic/gin"
)

// LoadApi add router for api service
func LoadApi(router *gin.Engine) {
	// DEMO for api data
	router.GET("/api/demo/get/:data", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"data": c.Param("data"),
		})
	})
	// TODO: edit to json for set, input json, add
	router.POST("/api/push/endpoint/add/", SetHandler)
	// TODO: edit to json for get, input string, obtain
	router.GET("/api/push/endpoint/list/", GetHandler)
}

func SetHandler(c *gin.Context){
	var endpnt push.Endpoint
	if err := c.ShouldBindJSON(&endpnt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,endpnt)
	// push.SetEndpoint(endpnt)
	return
}

func GetHandler(c *gin.Context){
	str:=c.Param("name")
	c.JSON(http.StatusOK,gin.H{"name":str})
	info:=push.GetEndpoint(str).Info()
	c.JSON(http.StatusOK,info)
	return
}