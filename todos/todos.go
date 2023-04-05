package todos

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Task        string    `json:"task" gorm:"type:varchar(255);not null"`
	CompletedAt time.Time `json:"completed_at" gorm:"type:timestamp"`
}
