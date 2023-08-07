package main

import (
	"github.com/gin-gonic/gin"

	"github.com/festivio/festivio-backend/config"
	"github.com/festivio/festivio-backend/internal/database"
	"github.com/festivio/festivio-backend/internal/server"
	"github.com/festivio/festivio-backend/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.GetLogger(cfg)

	db, err := database.NewPsqlDB(cfg)
	if err != nil {
		panic(err)
	}

	g := gin.New()
	s := server.NewServer(g, cfg, db, log)

	s.Run()
}