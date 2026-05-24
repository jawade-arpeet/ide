package v1

import (
	"ide/internal/handler"

	"github.com/gin-gonic/gin"
)

func mountHealthRoutes(
	router *gin.RouterGroup,
	handler *handler.HealthHandler,
) {
	router.GET("/health", handler.HealthCheck)
}
