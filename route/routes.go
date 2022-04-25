package route

import (
	"github.com/gin-gonic/gin"
	"websocket-boilerplate/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controller.SetupClientHtml)
	r.GET("/printLogs", controller.RunCommand)

	return r
}
