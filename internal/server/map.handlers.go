package server

import (
	"github.com/gin-gonic/gin"

	"github.com/festivio/festivio-backend/internal/handler"
	"github.com/festivio/festivio-backend/internal/repository"
	"github.com/festivio/festivio-backend/internal/service"
)

func (s *Server) MapHandlers(g *gin.Engine) {
	repo := repository.NewRepository(s.db, s.log)
	srv := service.NewService(repo, s.db, s.log)
	hdr := handler.NewHandler(srv, s.log)

	mainGroup := g.Group("/")
	handler.MapRoutes(mainGroup, hdr)
}