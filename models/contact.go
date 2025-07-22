package models

import (
	"time"

	"github.com/google/uuid"
)

// CREATE TYPE lead_status AS ENUM ('new', 'working', 'qualified', 'unqualified');
// CREATE TYPE lead_source AS ENUM (
//   'Website',
//   'Referral',
//   'LinkedIn',
//   'Facebook',
//   'Twitter',
//   'Instagram',
//   'YouTube',
//   'Google Ads',
//   'Facebook Ads',
//   'Instagram Ads',
//   'LinkedIn Ads',
//   'Email Campaign',
//   'Cold Call',
//   'Direct Visit',
//   'Webinar',
//   'Conference',
//   'Networking Event',
//   'Partnership',
//   'Reseller',
//   'Inbound',
//   'Outbound',
//   'Landing Page',
//   'Live Chat',
//   'Customer Support',
//   'SEO',
//   'Organic Search',
//   'Paid Search',
//   'Other'
// );

type Contact struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName   string
	LastName    string
	Email       string `gorm:"uniqueIndex"`
	Phone       string
	CompanyName string
	LeadStatus  string    `gorm:"type:lead_status"`
	Source      string    `gorm:"type:lead_source"`
	AssignedTo  uuid.UUID `gorm:"type:uuid"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
