package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Template statuses
const (
	TemplateStatusApproved = "approved"
	TemplateStatusPending  = "pending"
	TemplateStatusRejected = "rejected"
)

// Template categories
const (
	TemplateCategoryMarketing         = "marketing"
	TemplateCategoryUtility           = "utility"
	TemplateCategoryAuthentication    = "authentication"
)

// Template represents a WhatsApp message template
type Template struct {
	ID          string     `json:"id" gorm:"primaryKey;type:varchar(100)"`
	Name        string     `json:"name" gorm:"index;type:varchar(255);not null" validate:"required"`
	Language    string     `json:"language" gorm:"type:varchar(10);not null" validate:"required"`
	Category    string     `json:"category" gorm:"type:varchar(50);not null" validate:"required"`
	Status      string     `json:"status" gorm:"index;type:varchar(50);not null" validate:"required"`
	Content     string     `json:"content" gorm:"type:text;not null" validate:"required"`
	Parameters  JSONArray  `json:"parameters,omitempty" gorm:"type:jsonb"`
	Metadata    JSONMap    `json:"metadata,omitempty" gorm:"type:jsonb"`
	CreatedAt   time.Time  `json:"created_at" gorm:"index;not null"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"not null"`
}

// TableName specifies the table name for Template
func (Template) TableName() string {
	return "templates"
}

// BeforeCreate hook to generate ID and set timestamps
func (t *Template) BeforeCreate(tx *gorm.DB) error {
	if t.ID == "" {
		t.ID = GenerateID("tmpl")
	}
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now().UTC()
	}
	if t.UpdatedAt.IsZero() {
		t.UpdatedAt = time.Now().UTC()
	}
	return t.Validate()
}

// BeforeUpdate hook
func (t *Template) BeforeUpdate(tx *gorm.DB) error {
	t.UpdatedAt = time.Now().UTC()
	return nil
}

// Validate performs business logic validation
func (t *Template) Validate() error {
	if t.Name == "" {
		return errors.New("name is required")
	}
	if t.Language == "" {
		return errors.New("language is required")
	}
	if t.Category == "" {
		return errors.New("category is required")
	}
	if t.Status == "" {
		return errors.New("status is required")
	}
	if t.Content == "" {
		return errors.New("content is required")
	}

	// Validate category
	validCategories := []string{TemplateCategoryMarketing, TemplateCategoryUtility, TemplateCategoryAuthentication}
	if !contains(validCategories, t.Category) {
		return fmt.Errorf("invalid category: %s", t.Category)
	}

	// Validate status
	validStatuses := []string{TemplateStatusApproved, TemplateStatusPending, TemplateStatusRejected}
	if !contains(validStatuses, t.Status) {
		return fmt.Errorf("invalid status: %s", t.Status)
	}

	return nil
}

// IsApproved returns true if the template is approved
func (t *Template) IsApproved() bool {
	return t.Status == TemplateStatusApproved
}

// ParameterCount returns the number of parameters in the template
func (t *Template) ParameterCount() int {
	return len(t.Parameters)
}

// SubstituteParameters replaces placeholders with actual values
func (t *Template) SubstituteParameters(params []string) (string, error) {
	if len(params) != t.ParameterCount() {
		return "", fmt.Errorf("expected %d parameters, got %d", t.ParameterCount(), len(params))
	}

	result := t.Content
	for i, param := range params {
		placeholder := fmt.Sprintf("{{%d}}", i+1)
		result = strings.ReplaceAll(result, placeholder, param)
	}

	return result, nil
}

// Helper function to check if slice contains a value
func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
