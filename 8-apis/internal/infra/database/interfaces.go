package database

import "github.com/brenomachadodomonte/goexpert/apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
