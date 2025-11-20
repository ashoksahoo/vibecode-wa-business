# Data Models & Database Schema
# Vibecoded WA Client

**Last Updated:** November 18, 2025  
**Database:** PostgreSQL 15+

---

## Table of Contents

1. [Overview](#overview)
2. [Go Models](#go-models)
3. [Database Schema](#database-schema)
4. [Relationships](#relationships)
5. [Indexes](#indexes)
6. [Migrations](#migrations)

---

## Overview

The application uses PostgreSQL as the primary database with GORM as the ORM layer. All tables include `created_at` and `updated_at` timestamps for auditing.

### Design Principles
- **Normalized** structure for data integrity
- **JSONB** columns for flexible metadata
- **Indexes** on all foreign keys and query fields
- **UUID** strings for all primary keys
- **Soft deletes** where applicable

---

## Go Models

### Message Model

```go
package models

import (
    "time"
    "gorm.io/gorm"
)

type Message struct {
    ID              string         `json:"id" gorm:"primaryKey;type:varchar(255)"`
    Phone           string         `json:"phone" gorm:"type:varchar(20);not null;index"`
    Direction       string         `json:"direction" gorm:"type:varchar(10);not null"` // inbound, outbound
    Type            string         `json:"type" gorm:"type:varchar(20);not null"`      // text, image, document, template
    Content         string         `json:"content,omitempty" gorm:"type:text"`
    MediaURL        string         `json:"media_url,omitempty" gorm:"type:text"`
    Caption         string         `json:"caption,omitempty" gorm:"type:text"`
    Status          string         `json:"status" gorm:"type:varchar(20);not null;index"` // sent, delivered, read, failed
    WhatsAppMsgID   string         `json:"whatsapp_message_id,omitempty" gorm:"type:varchar(255);index"`
    ErrorCode       string         `json:"error_code,omitempty" gorm:"type:varchar(50)"`
    ErrorMessage    string         `json:"error_message,omitempty" gorm:"type:text"`
    Metadata        JSONMap        `json:"metadata,omitempty" gorm:"type:jsonb"`
    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
}

// Table name
func (Message) TableName() string {
    return "messages"
}

// Hooks
func (m *Message) BeforeCreate(tx *gorm.DB) error {
    if m.ID == "" {
        m.ID = generateID("msg")
    }
    return nil
}
```

### Contact Model

```go
package models

type Contact struct {
    ID                 string         `json:"id" gorm:"primaryKey;type:varchar(255)"`
    Phone              string         `json:"phone" gorm:"type:varchar(20);not null;uniqueIndex"`
    Name               string         `json:"name,omitempty" gorm:"type:varchar(255)"`
    ProfilePictureURL  string         `json:"profile_picture_url,omitempty" gorm:"type:text"`
    LastMessageAt      *time.Time     `json:"last_message_at,omitempty" gorm:"index"`
    MessageCount       int            `json:"message_count" gorm:"default:0"`
    UnreadCount        int            `json:"unread_count" gorm:"default:0"`
    Metadata           JSONMap        `json:"metadata,omitempty" gorm:"type:jsonb"`
    CreatedAt          time.Time      `json:"created_at"`
    UpdatedAt          time.Time      `json:"updated_at"`
    DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Contact) TableName() string {
    return "contacts"
}

func (c *Contact) BeforeCreate(tx *gorm.DB) error {
    if c.ID == "" {
        c.ID = generateID("contact")
    }
    return nil
}
```

### Template Model

```go
package models

type Template struct {
    ID          string         `json:"id" gorm:"primaryKey;type:varchar(255)"`
    Name        string         `json:"name" gorm:"type:varchar(255);not null;uniqueIndex"`
    Language    string         `json:"language" gorm:"type:varchar(10);not null"`
    Category    string         `json:"category" gorm:"type:varchar(50);not null"`
    Status      string         `json:"status" gorm:"type:varchar(20);not null"` // approved, pending, rejected
    Content     string         `json:"content" gorm:"type:text;not null"`
    Parameters  JSONArray      `json:"parameters,omitempty" gorm:"type:jsonb"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Template) TableName() string {
    return "templates"
}

func (t *Template) BeforeCreate(tx *gorm.DB) error {
    if t.ID == "" {
        t.ID = generateID("tmpl")
    }
    return nil
}
```

### APIKey Model

```go
package models

type APIKey struct {
    ID          string         `json:"id" gorm:"primaryKey;type:varchar(255)"`
    KeyHash     string         `json:"-" gorm:"type:varchar(255);not null;uniqueIndex"` // bcrypt hash
    Name        string         `json:"name" gorm:"type:varchar(255);not null"`
    Permissions JSONArray      `json:"permissions,omitempty" gorm:"type:jsonb"`
    RateLimit   int            `json:"rate_limit" gorm:"default:1000"`
    LastUsedAt  *time.Time     `json:"last_used_at,omitempty"`
    CreatedAt   time.Time      `json:"created_at"`
    ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
    DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

func (APIKey) TableName() string {
    return "api_keys"
}

func (k *APIKey) BeforeCreate(tx *gorm.DB) error {
    if k.ID == "" {
        k.ID = generateID("key")
    }
    return nil
}
```

### Custom Types

```go
package models

import (
    "database/sql/driver"
    "encoding/json"
    "errors"
)

// JSONMap for metadata fields
type JSONMap map[string]interface{}

func (j JSONMap) Value() (driver.Value, error) {
    return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
    if value == nil {
        *j = make(JSONMap)
        return nil
    }
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    return json.Unmarshal(bytes, j)
}

// JSONArray for arrays in JSONB
type JSONArray []string

func (j JSONArray) Value() (driver.Value, error) {
    return json.Marshal(j)
}

func (j *JSONArray) Scan(value interface{}) error {
    if value == nil {
        *j = []string{}
        return nil
    }
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    return json.Unmarshal(bytes, j)
}
```

---

## Database Schema

### Messages Table

```sql
CREATE TABLE messages (
    id VARCHAR(255) PRIMARY KEY,
    phone VARCHAR(20) NOT NULL,
    direction VARCHAR(10) NOT NULL CHECK (direction IN ('inbound', 'outbound')),
    type VARCHAR(20) NOT NULL CHECK (type IN ('text', 'image', 'document', 'video', 'audio', 'template')),
    content TEXT,
    media_url TEXT,
    caption TEXT,
    status VARCHAR(20) NOT NULL CHECK (status IN ('sent', 'delivered', 'read', 'failed')),
    whatsapp_msg_id VARCHAR(255),
    error_code VARCHAR(50),
    error_message TEXT,
    metadata JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_messages_phone ON messages(phone);
CREATE INDEX idx_messages_created_at ON messages(created_at DESC);
CREATE INDEX idx_messages_whatsapp_msg_id ON messages(whatsapp_msg_id);
CREATE INDEX idx_messages_status ON messages(status);
CREATE INDEX idx_messages_direction ON messages(direction);
CREATE INDEX idx_messages_type ON messages(type);

-- Full-text search index
CREATE INDEX idx_messages_content_fts ON messages USING gin(to_tsvector('english', content));

-- Composite indexes for common queries
CREATE INDEX idx_messages_phone_created ON messages(phone, created_at DESC);
CREATE INDEX idx_messages_status_created ON messages(status, created_at DESC);

-- Trigger for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_messages_updated_at BEFORE UPDATE ON messages
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

### Contacts Table

```sql
CREATE TABLE contacts (
    id VARCHAR(255) PRIMARY KEY,
    phone VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(255),
    profile_picture_url TEXT,
    last_message_at TIMESTAMP,
    message_count INTEGER DEFAULT 0,
    unread_count INTEGER DEFAULT 0,
    metadata JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Indexes
CREATE UNIQUE INDEX idx_contacts_phone ON contacts(phone) WHERE deleted_at IS NULL;
CREATE INDEX idx_contacts_last_message_at ON contacts(last_message_at DESC);
CREATE INDEX idx_contacts_name ON contacts(name);
CREATE INDEX idx_contacts_deleted_at ON contacts(deleted_at);

-- Full-text search index for name and phone
CREATE INDEX idx_contacts_search ON contacts USING gin(
    to_tsvector('english', COALESCE(name, '') || ' ' || phone)
);

-- Trigger for updated_at
CREATE TRIGGER update_contacts_updated_at BEFORE UPDATE ON contacts
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

### Templates Table

```sql
CREATE TABLE templates (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    language VARCHAR(10) NOT NULL,
    category VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('approved', 'pending', 'rejected')),
    content TEXT NOT NULL,
    parameters JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Indexes
CREATE UNIQUE INDEX idx_templates_name ON templates(name) WHERE deleted_at IS NULL;
CREATE INDEX idx_templates_status ON templates(status);
CREATE INDEX idx_templates_category ON templates(category);
CREATE INDEX idx_templates_deleted_at ON templates(deleted_at);

-- Trigger for updated_at
CREATE TRIGGER update_templates_updated_at BEFORE UPDATE ON templates
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
```

### API Keys Table

```sql
CREATE TABLE api_keys (
    id VARCHAR(255) PRIMARY KEY,
    key_hash VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    permissions JSONB,
    rate_limit INTEGER DEFAULT 1000,
    last_used_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Indexes
CREATE UNIQUE INDEX idx_api_keys_key_hash ON api_keys(key_hash) WHERE deleted_at IS NULL;
CREATE INDEX idx_api_keys_expires_at ON api_keys(expires_at);
CREATE INDEX idx_api_keys_deleted_at ON api_keys(deleted_at);
```

---

## Relationships

### Message ↔ Contact

```
Messages belong to Contacts (via phone number)
Contacts have many Messages
```

**Query Examples:**
```go
// Get all messages for a contact
var messages []Message
db.Where("phone = ?", contact.Phone).Find(&messages)

// Get contact for a message
var contact Contact
db.Where("phone = ?", message.Phone).First(&contact)
```

### Template → Messages

```
Messages can reference Templates (via metadata)
Templates have many Messages
```

**Query Example:**
```go
// Get all messages sent with a template
var messages []Message
db.Where("type = ? AND metadata->>'template_name' = ?", 
    "template", templateName).Find(&messages)
```

---

## Indexes

### Index Strategy

1. **Primary Keys**: All tables use VARCHAR primary keys
2. **Foreign Keys**: Indexed for join performance
3. **Frequent Queries**: Indexed phone, status, timestamps
4. **Full-Text Search**: GIN indexes for text search
5. **Composite Indexes**: For multi-column queries

### Index Monitoring

```sql
-- Check index usage
SELECT 
    schemaname,
    tablename,
    indexname,
    idx_scan as index_scans,
    idx_tup_read as tuples_read,
    idx_tup_fetch as tuples_fetched
FROM pg_stat_user_indexes
WHERE schemaname = 'public'
ORDER BY idx_scan DESC;

-- Find unused indexes
SELECT 
    schemaname,
    tablename,
    indexname
FROM pg_stat_user_indexes
WHERE idx_scan = 0
AND indexname NOT LIKE 'pg_toast%';
```

---

## Migrations

### Migration Strategy

1. **Version Control**: All migrations in `migrations/` directory
2. **Naming**: `YYYYMMDDHHMMSS_description.sql`
3. **Idempotent**: Can run multiple times safely
4. **Rollback**: Each migration has a down migration

### Example Migration

**File:** `migrations/20251118100000_create_messages_table.up.sql`
```sql
-- Create messages table
CREATE TABLE IF NOT EXISTS messages (
    id VARCHAR(255) PRIMARY KEY,
    phone VARCHAR(20) NOT NULL,
    direction VARCHAR(10) NOT NULL,
    type VARCHAR(20) NOT NULL,
    content TEXT,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_messages_phone ON messages(phone);
CREATE INDEX IF NOT EXISTS idx_messages_created_at ON messages(created_at DESC);
```

**File:** `migrations/20251118100000_create_messages_table.down.sql`
```sql
-- Drop indexes
DROP INDEX IF EXISTS idx_messages_phone;
DROP INDEX IF EXISTS idx_messages_created_at;

-- Drop table
DROP TABLE IF EXISTS messages;
```

### Running Migrations

```bash
# Using golang-migrate
migrate -path ./migrations -database "postgresql://user:pass@localhost:5432/wadb?sslmode=disable" up

# Using GORM AutoMigrate (development only)
db.AutoMigrate(&Message{}, &Contact{}, &Template{}, &APIKey{})
```

---

## Data Validation

### Model Validation

```go
package models

import "github.com/go-playground/validator/v10"

var validate = validator.New()

// Validate Message
func (m *Message) Validate() error {
    // Phone number validation
    if !isValidPhone(m.Phone) {
        return errors.New("invalid phone number format")
    }
    
    // Direction validation
    if m.Direction != "inbound" && m.Direction != "outbound" {
        return errors.New("direction must be inbound or outbound")
    }
    
    // Type validation
    validTypes := []string{"text", "image", "document", "video", "audio", "template"}
    if !contains(validTypes, m.Type) {
        return errors.New("invalid message type")
    }
    
    // Status validation
    validStatuses := []string{"sent", "delivered", "read", "failed"}
    if !contains(validStatuses, m.Status) {
        return errors.New("invalid status")
    }
    
    return nil
}

// Validate Contact
func (c *Contact) Validate() error {
    if !isValidPhone(c.Phone) {
        return errors.New("invalid phone number format")
    }
    return nil
}
```

---

## Query Examples

### Common Queries

```go
// Get recent messages for a phone number
var messages []Message
db.Where("phone = ?", phone).
   Order("created_at DESC").
   Limit(20).
   Find(&messages)

// Get undelivered messages
var messages []Message
db.Where("status IN ?", []string{"sent"}).
   Where("created_at < ?", time.Now().Add(-5*time.Minute)).
   Find(&messages)

// Search messages by content
var messages []Message
db.Where("content ILIKE ?", "%"+searchTerm+"%").
   Or("phone ILIKE ?", "%"+searchTerm+"%").
   Order("created_at DESC").
   Limit(50).
   Find(&messages)

// Get message statistics
type Stats struct {
    Total     int64
    Sent      int64
    Delivered int64
    Failed    int64
}

var stats Stats
db.Model(&Message{}).Count(&stats.Total)
db.Model(&Message{}).Where("status = ?", "sent").Count(&stats.Sent)
db.Model(&Message{}).Where("status = ?", "delivered").Count(&stats.Delivered)
db.Model(&Message{}).Where("status = ?", "failed").Count(&stats.Failed)

// Get or create contact
var contact Contact
result := db.Where("phone = ?", phone).FirstOrCreate(&contact, Contact{
    Phone: phone,
    Name:  extractNameFromMessage(message),
})

// Update contact last message time
db.Model(&Contact{}).
   Where("phone = ?", phone).
   Updates(map[string]interface{}{
       "last_message_at": time.Now(),
       "message_count":   gorm.Expr("message_count + ?", 1),
   })
```

---

## Performance Considerations

### Connection Pooling

```go
sqlDB, _ := db.DB()
sqlDB.SetMaxIdleConns(5)
sqlDB.SetMaxOpenConns(25)
sqlDB.SetConnMaxLifetime(time.Hour)
```

### Query Optimization

1. **Use Indexes**: Ensure all WHERE clauses use indexed columns
2. **Limit Results**: Always paginate large result sets
3. **Select Specific Columns**: Use `Select()` to fetch only needed fields
4. **Preload Relationships**: Use `Preload()` to avoid N+1 queries
5. **Use Raw SQL**: For complex queries, use raw SQL

### Monitoring Queries

```go
// Log slow queries (>100ms)
db.Logger = logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags),
    logger.Config{
        SlowThreshold: 100 * time.Millisecond,
        LogLevel:      logger.Warn,
    },
)
```

---

## Backup & Recovery

### Backup Strategy

```bash
# Daily backup
pg_dump -U postgres wadb > backup_$(date +%Y%m%d).sql

# Compressed backup
pg_dump -U postgres wadb | gzip > backup_$(date +%Y%m%d).sql.gz

# Backup specific tables
pg_dump -U postgres -t messages -t contacts wadb > backup_data.sql
```

### Restore

```bash
# Restore from backup
psql -U postgres wadb < backup_20251118.sql

# Restore compressed backup
gunzip -c backup_20251118.sql.gz | psql -U postgres wadb
```

---

## Database Maintenance

### Regular Maintenance Tasks

```sql
-- Vacuum tables
VACUUM ANALYZE messages;
VACUUM ANALYZE contacts;

-- Reindex tables
REINDEX TABLE messages;
REINDEX TABLE contacts;

-- Update statistics
ANALYZE messages;
ANALYZE contacts;
```

### Monitoring

```sql
-- Check table sizes
SELECT 
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) AS size
FROM pg_tables
WHERE schemaname = 'public'
ORDER BY pg_total_relation_size(schemaname||'.'||tablename) DESC;

-- Check database size
SELECT pg_size_pretty(pg_database_size('wadb'));
```

---

**Schema Version:** 1.0.0  
**Last Migration:** TBD
