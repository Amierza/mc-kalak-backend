package repository

import (
	"context"
	"errors"

	"github.com/Amierza/mc-kalak-backend/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	IVoteRepository interface {
		Create(ctx context.Context, tx *gorm.DB, vote *entity.Vote) error
		GetByClaimIDAndVoterID(ctx context.Context, tx *gorm.DB, claimID, voterID *uuid.UUID) (*entity.Vote, bool, error)
		GetAllByClaimID(ctx context.Context, tx *gorm.DB, claimID *uuid.UUID) ([]*entity.Vote, error)
	}

	voteRepository struct {
		db *gorm.DB
	}
)

func NewVoteRepository(db *gorm.DB) *voteRepository {
	return &voteRepository{
		db: db,
	}
}

func (vr *voteRepository) Create(ctx context.Context, tx *gorm.DB, vote *entity.Vote) error {
	if tx == nil {
		tx = vr.db
	}

	return tx.WithContext(ctx).Create(&vote).Error
}

func (vr *voteRepository) GetByClaimIDAndVoterID(ctx context.Context, tx *gorm.DB, claimID, voterID *uuid.UUID) (*entity.Vote, bool, error) {
	if tx == nil {
		tx = vr.db
	}

	var vote *entity.Vote
	err := tx.WithContext(ctx).
		Where("claim_id = ? AND voter_id = ?", &claimID, &voterID).
		Take(&vote).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &entity.Vote{}, false, nil
	}
	if err != nil {
		return &entity.Vote{}, false, err
	}

	return vote, true, nil
}

func (vr *voteRepository) GetAllByClaimID(ctx context.Context, tx *gorm.DB, claimID *uuid.UUID) ([]*entity.Vote, error) {
	if tx == nil {
		tx = vr.db
	}

	var (
		votes []*entity.Vote
		err   error
	)

	query := tx.WithContext(ctx).
		Preload("Claim.ClaimedPlayer").
		Preload("Claim.Reporter").
		Preload("Voter").
		Where("claim_id = ?", claimID).
		Model(&entity.Vote{})
	if err := query.Order(`"created_at" DESC`).Find(&votes).Error; err != nil {
		return []*entity.Vote{}, err
	}

	return votes, err
}
