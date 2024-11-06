package routers

import (
	"arcadia/middlewares"
	"io"

	"github.com/gin-gonic/gin"
)

// SetRouters sets the routers
func SetRouters() (routers *gin.Engine) {
	routers = gin.New()

	handlers := make([]gin.HandlerFunc, 0)
	handlers = append(handlers, middlewares.GenericRecovery())
	handlers = append(handlers, gin.Logger())

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	// 加载中间件
	routers.Use(handlers...)

	buildWebSocket(routers) // websocket

	return
}
