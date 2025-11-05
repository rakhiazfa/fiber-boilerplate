package entity

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"<-:create"`
	CreatedBy string    `gorm:"type:varchar(100)"`
	UpdatedAt time.Time `gorm:"<-"`
	UpdatedBy string    `gorm:"type:varchar(100)"`
}
