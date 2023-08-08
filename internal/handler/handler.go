package handler

import (
	"github.com/festivio/festivio-backend/config"
	"github.com/festivio/festivio-backend/internal/service"
	"github.com/festivio/festivio-backend/pkg/logger"
)

type handler struct {
	srv service.ServiceInterface
	log *logger.Logger
	cfg *config.Config
}

func NewHandler(srv service.ServiceInterface, log *logger.Logger, cfg *config.Config) HandlerInterface {
	return &handler{
		srv: srv,
		log: log,
		cfg: cfg,
	}
}
