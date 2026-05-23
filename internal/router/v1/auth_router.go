package v1

import (
	"ide/internal/handler"

	"github.com/gin-gonic/gin"
)

func mountAuthRoutes(
	router *gin.RouterGroup,
	handler *handler.AuthHandler,
) {
	authGroup := router.Group("/auth/")
	{
		authGroup.POST("/sign-in/", handler.SignIn)
	}

	profileGroup := router.Group("/profile/")
	{
		profileGroup.POST("/", handler.CreateProfile)
		profileGroup.GET("/", handler.GetProfile)
	}
}
