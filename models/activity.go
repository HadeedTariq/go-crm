package models

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Type        string
	Subject     string
	Description string
	ContactID   uuid.UUID `gorm:"type:uuid"`
	AssignedTo  uuid.UUID `gorm:"type:uuid"`
	DueDate     time.Time
	Completed   bool `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
