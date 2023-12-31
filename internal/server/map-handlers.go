package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/festivio/festivio-backend/docs"
	"github.com/festivio/festivio-backend/internal/handler"
	"github.com/festivio/festivio-backend/internal/repository"
	"github.com/festivio/festivio-backend/internal/service"
)

func (s *Server) MapHandlers(g *gin.Engine) {
	repo := repository.NewRepository(s.db, s.log)
	srv := service.NewService(repo, s.db, s.log)
	hdr := handler.NewHandler(srv, s.log, s.cfg)

	mainGroup := g.Group("/")
	mainGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	handler.MapRoutes(mainGroup, hdr, s.db)
}
