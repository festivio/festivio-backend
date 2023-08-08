package service

import (
	"errors"

	"github.com/festivio/festivio-backend/domain"
	"github.com/festivio/festivio-backend/pkg/utils"
)

func (s service) SignUpUser(signUpInput *domain.SignUpInput) error {
	hashedPassword, err := utils.HashPassword(signUpInput.Password)
	if err != nil {
		s.log.Err(err).Msg("password could not be hashed")
		return err
	}

	newUser := domain.User{
		Email:    signUpInput.Email,
		Name:     signUpInput.Name,
		Password: hashedPassword,
		Role:     signUpInput.Role,
	}

	err = s.repo.CreateUser(&newUser)
	if err != nil {
		s.log.Err(err).Msg("failed to create user")
		return err
	}

	return nil
}

func (s service) SignInUser(signInInput *domain.SignInInput) (*domain.User, error) {
	user, err := s.repo.GetUserByEmail(signInInput.Email)
	if err != nil {
		return nil, err
	}

	if err := utils.VerifyPassword(user.Password, signInInput.Password); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
