package service

import "github.com/festivio/festivio-backend/domain"

type ServiceInterface interface {
	// Authorization methods
	SignUpUser(signUpInput *domain.SignUpInput) error
	SignInUser(signInInput *domain.SignInInput) (*domain.User, error)
}
