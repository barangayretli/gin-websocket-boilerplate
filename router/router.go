package router

import (
	"net/http"

	"github.com/barangayretli/gin-websocket-boilerplate/handler"
	"github.com/barangayretli/gin-websocket-boilerplate/web"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", web.TestHTML)
	})
	r.GET("/ws", handler.RunCommand)
	r.GET("/health", handler.Health)

	return r
}
