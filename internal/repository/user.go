package repository

import (
	"errors"
	"strings"

	"github.com/festivio/festivio-backend/domain"
)

func (r repository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	tx := r.db.First(&user, "email = ?", strings.ToLower(email))
	if tx.Error != nil {
		r.log.Err(tx.Error).Msg("")
		return nil, errors.New("invalid email")
	}

	return &user, nil
}

func (r repository) GetUsers() ([]*domain.ShortUserInfo, error) {
	var users []*domain.ShortUserInfo
	tx := r.db.Table("users").Find(&users)
	if tx.Error != nil {
		r.log.Err(tx.Error).Msg("")
		return nil, tx.Error
	}

	return users, nil
}