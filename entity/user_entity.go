package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`

	ReportedClaims []Claim `gorm:"foreignKey:ReporterID;constraint:OnDelete:SET NULL;" json:"reported_claims,omitempty"`
	ClaimedEvents  []Claim `gorm:"foreignKey:ClaimedPlayerID;constraint:OnDelete:CASCADE;" json:"claimed_events,omitempty"`
	Votes          []Vote  `gorm:"foreignKey:VoterID;constraint:OnDelete:CASCADE;" json:"votes,omitempty"`

	TimeStamp
}
