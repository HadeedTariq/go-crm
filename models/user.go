package models

import (
	"time"

	"github.com/google/uuid"
)

// CREATE TYPE user_roles AS ENUM ('sales_rep', 'marketing_manager', 'customer_service_rep', 'admin');

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Role      string `gorm:"type:user_roles"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time

	Deals []Deal `gorm:"foreignKey:OwnerID"`
}
