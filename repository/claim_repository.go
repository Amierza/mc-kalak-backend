package repository

import (
	"context"
	"errors"
	"math"

	"github.com/Amierza/mc-kalak-backend/dto"
	"github.com/Amierza/mc-kalak-backend/entity"
	"github.com/Amierza/mc-kalak-backend/response"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	IClaimRepository interface {
		Create(ctx context.Context, tx *gorm.DB, claim *entity.Claim) error
		GetAllClaimsWithPagination(ctx context.Context, tx *gorm.DB, pagination response.PaginationRequest) (dto.ClaimPaginationRepositoryResponse, error)
		GetDetailByID(ctx context.Context, tx *gorm.DB, id *uuid.UUID) (*entity.Claim, bool, error)
		Update(ctx context.Context, tx *gorm.DB, claim *entity.Claim) error
		DeleteByID(ctx context.Context, tx *gorm.DB, id *uuid.UUID) error
	}

	claimRepository struct {
		db *gorm.DB
	}
)

func NewClaimRepository(db *gorm.DB) *claimRepository {
	return &claimRepository{
		db: db,
	}
}

func (cr *claimRepository) Create(ctx context.Context, tx *gorm.DB, claim *entity.Claim) error {
	if tx == nil {
		tx = cr.db
	}

	return tx.WithContext(ctx).Create(&claim).Error
}

func (cr *claimRepository) GetAllClaimsWithPagination(ctx context.Context, tx *gorm.DB, pagination response.PaginationRequest) (dto.ClaimPaginationRepositoryResponse, error) {
	if tx == nil {
		tx = cr.db
	}

	var (
		claims []*entity.Claim
		err    error
		count  int64
	)

	if pagination.PerPage == 0 {
		pagination.PerPage = 10
	}

	if pagination.Page == 0 {
		pagination.Page = 1
	}

	query := tx.WithContext(ctx).
		Preload("ClaimedPlayer").
		Preload("Reporter").
		Preload("Votes").
		Model(&entity.Claim{})

	if err := query.Order(`"created_at" DESC`).Find(&claims).Error; err != nil {
		return dto.ClaimPaginationRepositoryResponse{}, err
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.ClaimPaginationRepositoryResponse{}, err
	}

	if err := query.Scopes(response.Paginate(pagination.Page, pagination.PerPage)).Find(&claims).Error; err != nil {
		return dto.ClaimPaginationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(pagination.PerPage)))

	return dto.ClaimPaginationRepositoryResponse{
		Claims: claims,
		PaginationResponse: response.PaginationResponse{
			Page:    pagination.Page,
			PerPage: pagination.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}

func (cr *claimRepository) GetDetailByID(ctx context.Context, tx *gorm.DB, id *uuid.UUID) (*entity.Claim, bool, error) {
	if tx == nil {
		tx = cr.db
	}

	var claim *entity.Claim
	err := tx.WithContext(ctx).
		Preload("ClaimedPlayer").
		Preload("Reporter").
		Preload("Votes").
		Where("id = ?", &id).Take(&claim).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &entity.Claim{}, false, nil
	}
	if err != nil {
		return &entity.Claim{}, false, err
	}

	return claim, true, nil
}

func (cr *claimRepository) Update(ctx context.Context, tx *gorm.DB, claim *entity.Claim) error {
	if tx == nil {
		tx = cr.db
	}

	return tx.WithContext(ctx).Model(&entity.Claim{}).Where("id = ?", claim.ID).Updates(&claim).Error
}

func (cr *claimRepository) DeleteByID(ctx context.Context, tx *gorm.DB, id *uuid.UUID) error {
	if tx == nil {
		tx = cr.db
	}

	return tx.WithContext(ctx).Where("id = ?", id).Delete(&entity.Claim{}).Error
}
