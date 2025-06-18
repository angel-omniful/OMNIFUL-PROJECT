package models

import "time"

type InventoryAuditLog struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	InventoryID string    `gorm:"type:uuid;not null" json:"inventory_id"`
	Change      int       `gorm:"not null" json:"change"` // positive or negative
	Reason      string    `json:"reason"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
