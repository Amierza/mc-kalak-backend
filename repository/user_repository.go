package repository

import (
	"context"
	"errors"

	"github.com/Amierza/mc-kalak-backend/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	IUserRepository interface {
		Count(ctx context.Context, tx *gorm.DB) (int64, error)
		GetByUsername(ctx context.Context, tx *gorm.DB, username *string) (*entity.User, bool, error)
		GetDetailByID(ctx context.Context, tx *gorm.DB, id *uuid.UUID) (*entity.User, bool, error)
		Update(ctx context.Context, tx *gorm.DB, user *entity.User) error
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

func (ur *userRepository) Count(ctx context.Context, tx *gorm.DB) (int64, error) {
	if tx == nil {
		tx = ur.db
	}

	var count int64
	if err := tx.WithContext(ctx).Model(&entity.User{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (ur *userRepository) GetByUsername(ctx context.Context, tx *gorm.DB, username *string) (*entity.User, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var user *entity.User
	err := tx.WithContext(ctx).Where("username = ?", &username).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &entity.User{}, false, nil
	}
	if err != nil {
		return &entity.User{}, false, err
	}

	return user, true, nil
}

func (ur *userRepository) GetDetailByID(ctx context.Context, tx *gorm.DB, id *uuid.UUID) (*entity.User, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var user *entity.User
	err := tx.WithContext(ctx).Where("id = ?", &id).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &entity.User{}, false, nil
	}
	if err != nil {
		return &entity.User{}, false, err
	}

	return user, true, nil
}

func (ur *userRepository) Update(ctx context.Context, tx *gorm.DB, user *entity.User) error {
	if tx == nil {
		tx = ur.db
	}

	return tx.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID).Updates(&user).Error
}
