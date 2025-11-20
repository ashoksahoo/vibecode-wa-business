package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Message statuses
const (
	MessageStatusQueued    = "queued"
	MessageStatusSent      = "sent"
	MessageStatusDelivered = "delivered"
	MessageStatusRead      = "read"
	MessageStatusFailed    = "failed"
)

// Message types
const (
	MessageTypeText     = "text"
	MessageTypeImage    = "image"
	MessageTypeVideo    = "video"
	MessageTypeAudio    = "audio"
	MessageTypeDocument = "document"
	MessageTypeLocation = "location"
	MessageTypeTemplate = "template"
)

// Message represents a WhatsApp message
type Message struct {
	ID                  string    `json:"id" gorm:"primaryKey;type:varchar(100)"`
	WhatsAppMessageID   string    `json:"whatsapp_message_id" gorm:"uniqueIndex;type:varchar(255)"`
	FromNumber          string    `json:"from_number" gorm:"index;type:varchar(50);not null" validate:"required,e164"`
	ToNumber            string    `json:"to_number" gorm:"index;type:varchar(50);not null" validate:"required,e164"`
	Direction           string    `json:"direction" gorm:"type:varchar(20);not null" validate:"required,oneof=inbound outbound"`
	MessageType         string    `json:"message_type" gorm:"type:varchar(50);not null" validate:"required"`
	Content             string    `json:"content" gorm:"type:text"`
	MediaURL            string    `json:"media_url,omitempty" gorm:"type:varchar(500)"`
	MediaMimeType       string    `json:"media_mime_type,omitempty" gorm:"type:varchar(100)"`
	Status              string    `json:"status" gorm:"index;type:varchar(50);not null" validate:"required"`
	ErrorCode           string    `json:"error_code,omitempty" gorm:"type:varchar(100)"`
	ErrorMessage        string    `json:"error_message,omitempty" gorm:"type:text"`
	Metadata            JSONMap   `json:"metadata,omitempty" gorm:"type:jsonb"`
	Timestamp           time.Time `json:"timestamp" gorm:"index;not null"`
	CreatedAt           time.Time `json:"created_at" gorm:"index;not null"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"not null"`
}

// TableName specifies the table name for Message
func (Message) TableName() string {
	return "messages"
}

// BeforeCreate hook to generate ID and set timestamps
func (m *Message) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = GenerateID("msg")
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().UTC()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = time.Now().UTC()
	}
	if m.Timestamp.IsZero() {
		m.Timestamp = time.Now().UTC()
	}
	return m.Validate()
}

// BeforeUpdate hook to update timestamp
func (m *Message) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now().UTC()
	return nil
}

// Validate performs business logic validation
func (m *Message) Validate() error {
	if m.FromNumber == "" {
		return errors.New("from_number is required")
	}
	if m.ToNumber == "" {
		return errors.New("to_number is required")
	}
	if m.Direction != "inbound" && m.Direction != "outbound" {
		return fmt.Errorf("invalid direction: %s", m.Direction)
	}
	if m.MessageType == "" {
		return errors.New("message_type is required")
	}
	if m.Status == "" {
		return errors.New("status is required")
	}
	return nil
}

// IsInbound returns true if the message is inbound
func (m *Message) IsInbound() bool {
	return m.Direction == "inbound"
}

// IsOutbound returns true if the message is outbound
func (m *Message) IsOutbound() bool {
	return m.Direction == "outbound"
}

// IsDelivered returns true if the message has been delivered
func (m *Message) IsDelivered() bool {
	return m.Status == MessageStatusDelivered || m.Status == MessageStatusRead
}

// HasFailed returns true if the message has failed
func (m *Message) HasFailed() bool {
	return m.Status == MessageStatusFailed
}

// Contact represents a WhatsApp contact
type Contact struct {
	ID            string    `json:"id" gorm:"primaryKey;type:varchar(100)"`
	PhoneNumber   string    `json:"phone_number" gorm:"uniqueIndex;type:varchar(50);not null" validate:"required,e164"`
	Name          string    `json:"name" gorm:"type:varchar(255)"`
	ProfileURL    string    `json:"profile_url,omitempty" gorm:"type:varchar(500)"`
	LastMessageAt *time.Time `json:"last_message_at,omitempty" gorm:"index"`
	MessageCount  int       `json:"message_count" gorm:"default:0"`
	UnreadCount   int       `json:"unread_count" gorm:"default:0"`
	Metadata      JSONMap   `json:"metadata,omitempty" gorm:"type:jsonb"`
	CreatedAt     time.Time `json:"created_at" gorm:"index;not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"not null"`
}

// TableName specifies the table name for Contact
func (Contact) TableName() string {
	return "contacts"
}

// BeforeCreate hook to generate ID and set timestamps
func (c *Contact) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = GenerateID("contact")
	}
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now().UTC()
	}
	if c.UpdatedAt.IsZero() {
		c.UpdatedAt = time.Now().UTC()
	}
	return c.Validate()
}

// BeforeUpdate hook
func (c *Contact) BeforeUpdate(tx *gorm.DB) error {
	c.UpdatedAt = time.Now().UTC()
	return nil
}

// Validate performs business logic validation
func (c *Contact) Validate() error {
	if c.PhoneNumber == "" {
		return errors.New("phone_number is required")
	}
	return nil
}
