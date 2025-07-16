package models

import (
	"time"

	"github.com/google/uuid"
)

// CREATE TYPE lead_status AS ENUM ('new', 'working', 'qualified', 'unqualified');

type Contact struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName   string
	LastName    string
	Email       string `gorm:"uniqueIndex"`
	Phone       string
	CompanyName string
	LeadStatus  string `gorm:"type:lead_status"`
	Source      string
	AssignedTo  uuid.UUID `gorm:"type:uuid"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
