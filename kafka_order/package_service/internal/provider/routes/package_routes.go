package routes

import (
	"package_service/internal/handler"

	"github.com/gin-gonic/gin"
)

func PackageRoutes(routerGroup *gin.RouterGroup, userHandler handler.PackageHandlerInterface) {
	routerGroup.GET("/", userHandler.GetAllOrFindByName)
	routerGroup.GET("/:id", userHandler.FindById)
}
