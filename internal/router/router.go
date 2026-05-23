package router

import (
	"ide/internal/handler"
	v1 "ide/internal/router/v1"
	"ide/internal/types"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	runEnv types.Env,
	handler *handler.Handler,
) *gin.Engine {
	if runEnv == types.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	v1.MountV1Routes(router, handler)

	return router
}
