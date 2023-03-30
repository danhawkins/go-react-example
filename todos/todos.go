package todos

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Task        string
	CompletedAt time.Time
}

func Setup(db *gorm.DB) error {
	return db.AutoMigrate(&Todo{})
}
