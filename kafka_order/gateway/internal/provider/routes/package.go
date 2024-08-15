package routes

import (
	"gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func PackageRoutes(routerGroup *gin.RouterGroup, host string) {
	routerGroup.GET("/", middleware.PseudoHandler(host))
	routerGroup.GET("/:id", middleware.PseudoHandler(host))
}
