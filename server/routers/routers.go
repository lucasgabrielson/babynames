package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"./handlers/handle"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello,", handle.HalloHandler)
	return r
}
