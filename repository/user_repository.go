package repository

import (
	"gorm.io/gorm"
)

type (
	IUserRepository interface {
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}
