package service

import "github.com/festivio/festivio-backend/domain"

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type ServiceInterface interface {
	// Authorization methods
	SignUpUser(signUpInput *domain.SignUpInput) error
	SignInUser(signInInput *domain.SignInInput) (*domain.User, error)
}
