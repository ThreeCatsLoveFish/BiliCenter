package api

import (
	"github.com/gin-gonic/gin"
)

type StatusCode int32

// Status code
const (
	Success StatusCode = iota
	ErrParam
	ErrInternal
)

type Resp struct {
	Code StatusCode  `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// LoadApi add router for api service
func LoadApi(router *gin.Engine) {
	// Push
	router.GET("/api/push/endpoint/list/", ListEndpoint)
	router.POST("/api/push/endpoint/update/", UpdateEndpoint)
	// Bili account
	router.GET("/api/bili/update/", UpdateBili)
}
