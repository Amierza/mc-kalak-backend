package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Claim struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	Event         ClaimEvent  `gorm:"type:varchar(10);not null" json:"event"`
	Status        ClaimStatus `gorm:"type:varchar(20);default:PENDING" json:"status"`
	MatchDate     time.Time   `gorm:"type:date;not null;index" json:"match_date"`
	TotalPlayer   int         `gorm:"not null" json:"total_player"`
	ScreenshotURL string      `gorm:"not null" json:"screenshot_url"`

	ApproveCount int `gorm:"default:0" json:"approve_count"`
	RejectCount  int `gorm:"default:0" json:"reject_count"`

	ClaimedPlayerID uuid.UUID `gorm:"type:uuid;index;not null" json:"claimed_player_id"`
	ClaimedPlayer   User      `gorm:"foreignKey:ClaimedPlayerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"claimed_player"`

	ReporterID uuid.UUID `gorm:"type:uuid;index;not null" json:"reporter_id"`
	Reporter   User      `gorm:"foreignKey:ReporterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reporter"`

	Votes []Vote `gorm:"foreignKey:ClaimID;constraint:OnDelete:CASCADE;" json:"votes,omitempty"`

	TimeStamp
}

func (c *Claim) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()

	if c.TotalPlayer <= 0 {
		return err
	}

	return nil
}
