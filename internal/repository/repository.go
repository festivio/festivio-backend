package repository

import (
	"gorm.io/gorm"

	"github.com/festivio/festivio-backend/pkg/logger"
)

type repository struct {
	db  *gorm.DB
	log *logger.Logger
}

func NewRepository(db *gorm.DB, log *logger.Logger) RepositoryInterface {
	return &repository{
		db:  db,
		log: log,
	}
}
