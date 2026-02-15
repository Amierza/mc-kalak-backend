package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vote struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	ClaimID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_claim_voter" json:"claim_id"`
	Claim   Claim     `gorm:"foreignKey:ClaimID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"claim"`

	VoterID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_claim_voter" json:"voter_id"`
	Voter   User      `gorm:"foreignKey:VoterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"voter"`

	Type VoteType `gorm:"type:varchar(10);not null" json:"type"`

	TimeStamp
}

func (v *Vote) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.New()
	return
}
