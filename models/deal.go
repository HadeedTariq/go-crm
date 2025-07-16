package models

import (
	"time"

	"github.com/google/uuid"
)

// CREATE TYPE deal_stage AS ENUM ('prospecting', 'negotiation', 'closed_won', 'closed_lost');

type Deal struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string
	ContactID uuid.UUID `gorm:"type:uuid"`
	Contact   Contact   `gorm:"foreignKey:ContactID"`

	Value     float64
	Stage     string `gorm:"type:deal_stage"`
	CloseDate time.Time

	OwnerID uuid.UUID `gorm:"type:uuid"`
	Owner   User      `gorm:"foreignKey:OwnerID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
