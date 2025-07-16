package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string
	Industry  string
	Size      string
	Website   string
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
