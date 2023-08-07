package service

import (
	"gorm.io/gorm"

	"github.com/festivio/festivio-backend/internal/repository"
	"github.com/festivio/festivio-backend/pkg/logger"
)

type service struct {
	repo repository.RepositoryInterface
	db   *gorm.DB
	log  *logger.Logger
}

func NewService(repo repository.RepositoryInterface, db *gorm.DB, log *logger.Logger) ServiceInterface {
	return &service{
		repo: repo,
		db:   db,
		log:  log,
	}
}
