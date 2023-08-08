package repository

import "github.com/festivio/festivio-backend/domain"

type RepositoryInterface interface {
	// User methods
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
}
