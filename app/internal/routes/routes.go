package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/nannigalaxy/go-rest-api/app/internal/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	api_router := router.Group("/api/v1")
	controllers.UserRouter(api_router)
}
