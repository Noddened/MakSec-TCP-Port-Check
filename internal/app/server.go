package app

import (
	"log/slog"

	"github.com/Noddened/MakSec-TCP-Port-Check/internal/api"
	"github.com/Noddened/MakSec-TCP-Port-Check/internal/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
		router: gin.Default(),
	}
}

func (s *Server) Start(logger *slog.Logger) error {
	err := setupRouter(s)
	if err != nil {
		logger.Error("Ошибка настройки роутера", "error", err)
		return err
	}

	return nil
}

func setupRouter(s *Server) error {
	s.router.POST("/check", api.HandleCheck)

	s.router.SetTrustedProxies([]string{"127.0.0.1"})
	gin.SetMode(gin.ReleaseMode)
	s.router.Run(s.config.BindAddr)
	return nil
}
