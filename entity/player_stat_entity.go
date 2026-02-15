package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlayerStat struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	TotalMatch int `gorm:"default:0" json:"total_match"`
	KingCount  int `gorm:"default:0" json:"king_count"`
	KongCount  int `gorm:"default:0" json:"kong_count"`
	NgokCount  int `gorm:"default:0" json:"ngok_count"`

	Score   int     `gorm:"default:0" json:"score"`
	WinRate float64 `gorm:"default:0" json:"win_rate"`

	PlayerID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"player_id"`
	Player   User      `gorm:"foreignKey:PlayerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"player"`

	TimeStamp
}

func (ps *PlayerStat) BeforeCreate(tx *gorm.DB) (err error) {
	ps.ID = uuid.New()
	return
}
