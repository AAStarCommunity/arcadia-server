package routers

import (
	"arcadia/controllers"

	"github.com/gin-gonic/gin"
)

func buildWebSocket(router *gin.Engine) {
	router.GET(controllers.ROUTER_HEALTHZ, controllers.Healthz)
}
