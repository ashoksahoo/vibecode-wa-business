package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel contains common fields for all models
type BaseModel struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(100)"`
	CreatedAt time.Time      `json:"created_at" gorm:"index;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// BeforeCreate hook to generate UUID-based ID
func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	if b.CreatedAt.IsZero() {
		b.CreatedAt = time.Now().UTC()
	}
	if b.UpdatedAt.IsZero() {
		b.UpdatedAt = time.Now().UTC()
	}
	return nil
}

// BeforeUpdate hook to update the UpdatedAt timestamp
func (b *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now().UTC()
	return nil
}

// GenerateID generates a prefixed ID (e.g., "msg_abc123")
func GenerateID(prefix string) string {
	id := uuid.New().String()
	if prefix != "" {
		return prefix + "_" + id
	}
	return id
}
