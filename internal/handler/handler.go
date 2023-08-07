package handler

import (
	"github.com/festivio/festivio-backend/internal/service"
	"github.com/festivio/festivio-backend/pkg/logger"
)

type handler struct {
	srv service.ServiceInterface
	log *logger.Logger
}

func NewHandler(srv service.ServiceInterface, log *logger.Logger) HandlerInterface {
	return &handler{
		srv: srv,
		log: log,
	}
}
