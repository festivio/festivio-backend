package repository

import (
	"errors"
	"strings"

	"github.com/festivio/festivio-backend/domain"
)

func (r repository) CreateUser(user *domain.User) error {
	tx := r.db.Create(&user)
	if tx.Error != nil && strings.Contains(tx.Error.Error(), "duplicate key value violates unique") {
		if strings.Contains(tx.Error.Error(), `"idx_users_email"`) {
			return errors.New("user with this email address already exists")
		}
		return errors.New("is non-unique data")
	} else if tx.Error != nil {
		r.log.Err(tx.Error).Msg("")
		return errors.New("something bad happened")
	}

	return nil
}

func (r repository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	tx := r.db.First(&user, "email = ?", strings.ToLower(email))
	if tx.Error != nil {
		r.log.Err(tx.Error).Msg("")
		return nil, errors.New("invalid email")
	}

	return &user, nil
}
