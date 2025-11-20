package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// APIKey represents an API key for authentication
type APIKey struct {
	ID          string     `json:"id" gorm:"primaryKey;type:varchar(100)"`
	Name        string     `json:"name" gorm:"type:varchar(255);not null" validate:"required"`
	KeyHash     string     `json:"-" gorm:"uniqueIndex;type:varchar(255);not null"`
	KeyPrefix   string     `json:"key_prefix" gorm:"index;type:varchar(20)"`
	Permissions JSONArray  `json:"permissions,omitempty" gorm:"type:jsonb"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty" gorm:"index"`
	LastUsedAt  *time.Time `json:"last_used_at,omitempty" gorm:"index"`
	CreatedAt   time.Time  `json:"created_at" gorm:"index;not null"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"not null"`
}

// TableName specifies the table name for APIKey
func (APIKey) TableName() string {
	return "api_keys"
}

// BeforeCreate hook to generate ID and set timestamps
func (a *APIKey) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = GenerateID("key")
	}
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now().UTC()
	}
	if a.UpdatedAt.IsZero() {
		a.UpdatedAt = time.Now().UTC()
	}
	return a.Validate()
}

// BeforeUpdate hook
func (a *APIKey) BeforeUpdate(tx *gorm.DB) error {
	a.UpdatedAt = time.Now().UTC()
	return nil
}

// Validate performs business logic validation
func (a *APIKey) Validate() error {
	if a.Name == "" {
		return errors.New("name is required")
	}
	if a.KeyHash == "" {
		return errors.New("key_hash is required")
	}
	return nil
}

// IsExpired returns true if the API key has expired
func (a *APIKey) IsExpired() bool {
	if a.ExpiresAt == nil {
		return false
	}
	return time.Now().UTC().After(*a.ExpiresAt)
}

// IsValid returns true if the API key is valid (not expired)
func (a *APIKey) IsValid() bool {
	return !a.IsExpired()
}

// UpdateLastUsed updates the last_used_at timestamp
func (a *APIKey) UpdateLastUsed(tx *gorm.DB) error {
	now := time.Now().UTC()
	a.LastUsedAt = &now
	return tx.Model(a).Update("last_used_at", now).Error
}

// HasPermission checks if the API key has a specific permission
func (a *APIKey) HasPermission(permission string) bool {
	if a.Permissions == nil {
		return false
	}
	for _, p := range a.Permissions {
		if p == permission || p == "*" {
			return true
		}
	}
	return false
}
