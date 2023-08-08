package main

import (
	"github.com/gin-gonic/gin"

	"github.com/festivio/festivio-backend/config"
	"github.com/festivio/festivio-backend/internal/database"
	"github.com/festivio/festivio-backend/internal/server"
	"github.com/festivio/festivio-backend/pkg/logger"
)

// @title           Festivio API
// @version         1.0
// @description     REST API for Festivio App

// @contact.url    https://github.com/festivio/festivio-backend

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
