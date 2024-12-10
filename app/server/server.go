package server

import (
	"github.com/nannigalaxy/go-rest-api/app/connections/logging"
	"github.com/nannigalaxy/go-rest-api/app/internal/routes"

	"github.com/gin-gonic/gin"

	"github.com/rollbar/rollbar-go"
)

func Initialize() {
	logging.SetRollbarConfig()
	defer rollbar.Close()
	app_router := gin.Default()
	app_router.Use(logging.RollbarMiddleware())
	routes.RegisterRoutes(app_router)
	app_router.Run("localhost:3000")
}
