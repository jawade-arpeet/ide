package v1

import (
	"ide/internal/handler"
	"ide/internal/middleware"

	"github.com/gin-gonic/gin"
)

func MountV1Routes(
	router *gin.Engine,
	middleware *middleware.Middleware,
	handler *handler.Handler,
) {
	v1 := router.Group("/api/v1")

	mountHealthRoutes(v1, middleware, handler.Health)
	mountAuthRoutes(v1, handler.Auth)
}
