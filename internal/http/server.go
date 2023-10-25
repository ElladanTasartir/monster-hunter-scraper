package http

import (
	"fmt"
	"github.com/ElladanTasartir/monster-hunter-scraper/internal/config"
	"github.com/ElladanTasartir/monster-hunter-scraper/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	httpServer *gin.Engine
	storage    *storage.Storage
	address    string
}

func NewServer(appConfig *config.Config, database *storage.Storage) (*Server, error) {
	engine := gin.Default()

	server := &Server{
		httpServer: engine,
		address:    fmt.Sprintf(":%d", appConfig.Port),
		storage:    database,
	}

	server.httpServer.NoRoute(server.NotFound)
	server.httpServer.GET("/", server.Healthcheck)

	err := server.NewScraperEndpoint(appConfig.ApiURL)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func (s *Server) Run() error {
	err := s.httpServer.Run(s.address)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Resource not found",
	})
}

func (s *Server) Healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})
}
