package v1

import (
	"ide/internal/handler"
	"ide/internal/middleware"

	"github.com/gin-gonic/gin"
)

func mountWorkspaceRoutes(
	router *gin.RouterGroup,
	middleware *middleware.Middleware,
	handler *handler.WorkspaceHandler,
) {
	workspaceGroup := router.Group(
		"/workspace",
		middleware.AuthMiddleware(),
	)

	{
		workspaceGroup.POST("", handler.CreateWorkspace)
	}
}
