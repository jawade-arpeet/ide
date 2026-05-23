package server

import (
	"fmt"
	"ide/internal/client"
	"ide/internal/config"
	"ide/internal/handler"
	"ide/internal/middleware"
	"ide/internal/repository"
	"ide/internal/router"
	"ide/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	config *config.ServerConfig
	router *gin.Engine
}

func NewServer(
	serverCfg *config.ServerConfig,
	client *client.Client,
) *Server {
	repo := repository.NewRepository(client)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	middleware := middleware.NewMiddleware()
	router := router.NewRouter(serverCfg.RunEnv, middleware, handler)

	return &Server{
		config: serverCfg,
		router: router,
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.config.Port)

	zap.L().Info("starting server",
		zap.String("address", addr),
	)

	return s.router.Run(addr)
}
