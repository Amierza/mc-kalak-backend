package migrations

import (
	"github.com/Amierza/mc-kalak-backend/entity"
	"gorm.io/gorm"
)

func Rollback(db *gorm.DB) error {
	tables := []interface{}{
		&entity.PlayerStat{},
		&entity.Vote{},
		&entity.Claim{},
		&entity.User{},
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	return nil
}
