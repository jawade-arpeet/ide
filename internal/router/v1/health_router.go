package v1

import (
	"ide/internal/handler"
	"ide/internal/middleware"

	"github.com/gin-gonic/gin"
)

func mountHealthRoutes(
	router *gin.RouterGroup,
	middleware *middleware.Middleware,
	handler *handler.HealthHandler,
) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/health", handler.HealthCheck)
}
