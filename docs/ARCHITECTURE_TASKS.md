# Architecture Implementation Tasks
# Vibecoded WA Client

**Last Updated:** November 21, 2025
**Status:** 🟢 Core Foundation Complete

## ✅ Completed Tasks

The following architecture tasks have been implemented:

1. **Project Structure Setup** ✅
   - Directory structure with all required packages
   - Go module initialization
   - All dependencies added

2. **Database Layer** ✅
   - Connection management with GORM
   - Health check system
   - Auto-migration support
   - Custom indexes and triggers

3. **Data Models** ✅
   - Message, Contact, Template, APIKey models
   - Call, Transcript models (for Phase 4)
   - Custom JSONB types
   - Model validation and hooks

4. **Configuration Management** ✅
   - Viper-based configuration
   - Environment variable support
   - Validation and defaults

5. **Logging System** ✅
   - Zap structured logging
   - Multiple output formats
   - Request ID tracking

6. **Error Handling** ✅
   - Custom error types with codes
   - HTTP status mapping
   - Detailed error responses

7. **Utilities** ✅
   - Pagination helpers
   - Response helpers
   - Crypto utilities (API key generation, hashing)
   - Input validators

8. **Main Application** ✅
   - Graceful shutdown
   - Database initialization
   - Signal handling

---

## Overview

This document breaks down the architecture decisions and system design into actionable implementation tasks. Each section corresponds to architectural decisions and includes clear TODO markers for Claude Code.

---

## Table of Contents

1. [Project Structure Setup](#1-project-structure-setup)
2. [Database Layer](#2-database-layer)
3. [WhatsApp API Integration](#3-whatsapp-api-integration)
4. [API Server Setup](#4-api-server-setup)
5. [Authentication & Security](#5-authentication--security)
6. [Service Layer](#6-service-layer)
7. [Repository Layer](#7-repository-layer)
8. [Middleware Stack](#8-middleware-stack)
9. [Configuration Management](#9-configuration-management)
10. [Error Handling](#10-error-handling)
11. [Logging System](#11-logging-system)
12. [Docker & Deployment](#12-docker--deployment)

---

## 1. Project Structure Setup

### Task 1.1: Initialize Go Module
**Architecture Decision:** Go 1.21+ with Go modules

**Subtasks:**
- [x] Create project root directory structure
- [x] Initialize Go module with proper naming
- [x] Set up directory structure following Go conventions

**Directory Structure Needed:**
```
vibecoded-wa-client/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── routes/
│   ├── config/
│   ├── database/
│   ├── models/
│   ├── repositories/
│   ├── services/
│   └── whatsapp/
├── pkg/
│   ├── errors/
│   └── utils/
├── migrations/
├── docs/
├── scripts/
├── .env.example
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
├── docker-compose.yml
└── README.md
```

**TODO: CLAUDE_CODE**
- Create complete directory structure
- Initialize go.mod with project path
- Create placeholder files in each directory
- Add README files explaining each directory's purpose

---

### Task 1.2: Dependency Management
**Architecture Decision:** Go modules with explicit versioning

**Dependencies to Add:**
```go
// Web Framework
github.com/gin-gonic/gin v1.9.1

// Database
gorm.io/gorm v1.25.5
gorm.io/driver/postgres v1.5.4

// HTTP Client
github.com/go-resty/resty/v2 v2.10.0

// Configuration
github.com/spf13/viper v1.18.0
github.com/joho/godotenv v1.5.1

// Logging
go.uber.org/zap v1.26.0

// Validation
github.com/go-playground/validator/v10 v10.16.0

// UUID
github.com/google/uuid v1.4.0

// Testing
github.com/stretchr/testify v1.8.4

// Metrics (Optional)
github.com/prometheus/client_golang v1.17.0

// WebSocket (for future WebRTC)
github.com/gorilla/websocket v1.5.1

// Rate Limiting
golang.org/x/time v0.5.0
```

**TODO: CLAUDE_CODE**
- Add all dependencies to go.mod
- Run `go mod tidy`
- Create vendor directory if needed
- Document dependency versions in DEPENDENCIES.md

---

## 2. Database Layer

### Task 2.1: Database Connection Setup
**Architecture Decision:** PostgreSQL with GORM ORM

**Implementation Steps:**
- [x] Create database connection configuration
- [x] Implement connection pool settings
- [x] Add connection retry logic
- [x] Create health check for database

**Files to Create:**
- `internal/database/connection.go`
- `internal/database/health.go`

**TODO: CLAUDE_CODE**
```go
// internal/database/connection.go
package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

// TODO: CLAUDE_CODE
// Implement:
// - NewConnection(dsn string) (*gorm.DB, error)
// - SetConnectionPool(db *gorm.DB, config PoolConfig)
// - CloseConnection(db *gorm.DB) error
// - PingDatabase(db *gorm.DB) error

// Configuration:
// - MaxIdleConns: 5
// - MaxOpenConns: 25
// - ConnMaxLifetime: 1 hour
```

---

### Task 2.2: Database Models
**Architecture Decision:** GORM models with proper tags and validation

**Models to Implement:**
1. Message
2. Contact
3. Template
4. APIKey
5. Call (future)
6. Transcription (future)

**Files to Create:**
- `internal/models/message.go`
- `internal/models/contact.go`
- `internal/models/template.go`
- `internal/models/api_key.go`
- `internal/models/types.go` (for custom types like JSONMap)
- `internal/models/base.go` (for common fields/methods)

**TODO: CLAUDE_CODE**
```go
// internal/models/message.go
package models

// TODO: CLAUDE_CODE
// Implement Message model with:
// - All fields from DATA_MODELS.md
// - GORM tags (primaryKey, index, type, etc.)
// - JSON tags for API responses
// - Validation tags
// - TableName() method
// - BeforeCreate() hook for ID generation
// - Validate() method for business logic validation

// internal/models/contact.go
// TODO: CLAUDE_CODE
// Implement Contact model with similar structure

// internal/models/template.go
// TODO: CLAUDE_CODE
// Implement Template model with similar structure

// internal/models/api_key.go
// TODO: CLAUDE_CODE
// Implement APIKey model with similar structure

// internal/models/types.go
// TODO: CLAUDE_CODE
// Implement custom types:
// - JSONMap (map[string]interface{})
// - JSONArray ([]string)
// Both with Value() and Scan() methods for GORM
```

---

### Task 2.3: Database Migrations
**Architecture Decision:** SQL migrations with up/down files

**Migration Files Needed:**
1. `001_create_messages_table.up.sql`
2. `001_create_messages_table.down.sql`
3. `002_create_contacts_table.up.sql`
4. `002_create_contacts_table.down.sql`
5. `003_create_templates_table.up.sql`
6. `003_create_templates_table.down.sql`
7. `004_create_api_keys_table.up.sql`
8. `004_create_api_keys_table.down.sql`

**TODO: CLAUDE_CODE**
- Create all migration files in `migrations/` directory
- Copy SQL from DATA_MODELS.md
- Add proper indexes as specified
- Include trigger creation for updated_at
- Create migration runner script

**Migration Runner:**
```go
// internal/database/migrate.go
package database

// TODO: CLAUDE_CODE
// Implement:
// - RunMigrations(db *gorm.DB, migrationsPath string) error
// - RollbackMigration(db *gorm.DB, migrationsPath string) error
// - GetMigrationStatus(db *gorm.DB) ([]Migration, error)
```

---

## 3. WhatsApp API Integration

### Task 3.1: WhatsApp Client Wrapper
**Architecture Decision:** WhatsApp Cloud API with abstraction layer

**Implementation Structure:**
- [ ] Create WhatsApp client struct
- [ ] Implement authentication
- [ ] Add request/response models
- [ ] Implement error handling
- [ ] Add retry logic

**Files to Create:**
- `internal/whatsapp/client.go`
- `internal/whatsapp/messages.go`
- `internal/whatsapp/webhooks.go`
- `internal/whatsapp/models.go`
- `internal/whatsapp/errors.go`

**TODO: CLAUDE_CODE**
```go
// internal/whatsapp/client.go
package whatsapp

import "github.com/go-resty/resty/v2"

// TODO: CLAUDE_CODE
// Implement:
// - WhatsAppClient struct with:
//   - httpClient *resty.Client
//   - apiToken string
//   - phoneNumberID string
//   - baseURL string
// - NewClient(config Config) (*WhatsAppClient, error)
// - SetTimeout(duration time.Duration)
// - SetRetryPolicy(count int, waitTime time.Duration)

// internal/whatsapp/messages.go
// TODO: CLAUDE_CODE
// Implement message sending methods:
// - SendTextMessage(to, text string) (*MessageResponse, error)
// - SendMediaMessage(to, mediaURL, caption string, mediaType MediaType) (*MessageResponse, error)
// - SendTemplateMessage(to, templateName, language string, params []string) (*MessageResponse, error)
// - GetMessageStatus(messageID string) (*MessageStatus, error)

// internal/whatsapp/webhooks.go
// TODO: CLAUDE_CODE
// Implement webhook handling:
// - VerifySignature(body []byte, signature string, secret string) bool
// - ParseWebhook(body []byte) (*WebhookPayload, error)
// - ParseMessageEvent(payload *WebhookPayload) (*MessageEvent, error)
// - ParseStatusEvent(payload *WebhookPayload) (*StatusEvent, error)

// internal/whatsapp/models.go
// TODO: CLAUDE_CODE
// Define all request/response models:
// - MessageRequest, MessageResponse
// - WebhookPayload, MessageEvent, StatusEvent
// - MediaType constants
// - Error models

// internal/whatsapp/errors.go
// TODO: CLAUDE_CODE
// Define error types:
// - WhatsAppError struct
// - Error code constants
// - Error mapping from API responses
```

---

## 4. API Server Setup

### Task 4.1: Gin Server Initialization
**Architecture Decision:** Gin web framework with middleware pipeline

**Implementation Steps:**
- [ ] Create server struct
- [ ] Initialize Gin engine
- [ ] Configure Gin mode (debug/release)
- [ ] Set up graceful shutdown
- [ ] Add signal handling

**Files to Create:**
- `internal/api/server.go`
- `internal/api/routes/routes.go`
- `cmd/server/main.go`

**TODO: CLAUDE_CODE**
```go
// internal/api/server.go
package api

import "github.com/gin-gonic/gin"

// TODO: CLAUDE_CODE
// Implement:
// - Server struct with:
//   - router *gin.Engine
//   - config *config.Config
//   - db *gorm.DB
//   - whatsappClient *whatsapp.WhatsAppClient
//   - services (initialized later)
// - NewServer(config *config.Config) (*Server, error)
// - SetupRoutes()
// - SetupMiddleware()
// - Start() error
// - Shutdown(ctx context.Context) error
// - GracefulShutdown() with signal handling

// cmd/server/main.go
// TODO: CLAUDE_CODE
// Implement main entry point:
// - Load configuration
// - Initialize database
// - Initialize WhatsApp client
// - Create and start server
// - Handle shutdown signals
// - Cleanup resources
```

---

### Task 4.2: Route Registration
**Architecture Decision:** Organized route groups by resource

**Route Groups:**
1. `/api/v1/messages` - Message endpoints
2. `/api/v1/contacts` - Contact endpoints
3. `/api/v1/templates` - Template endpoints
4. `/api/v1/calls` - Call endpoints (future)
5. `/webhooks/whatsapp` - Webhook endpoint
6. `/health` - Health check
7. `/metrics` - Metrics endpoint

**TODO: CLAUDE_CODE**
```go
// internal/api/routes/routes.go
package routes

// TODO: CLAUDE_CODE
// Implement:
// - RegisterRoutes(router *gin.Engine, handlers *Handlers) where Handlers contains all handler instances
// - Create route groups with proper middleware
// - Register all endpoints from API_DESIGN.md

// Structure:
// v1 := router.Group("/api/v1")
// v1.Use(AuthMiddleware())
// {
//     messages := v1.Group("/messages")
//     {
//         messages.POST("", handlers.SendMessage)
//         messages.GET("", handlers.ListMessages)
//         messages.GET("/:id", handlers.GetMessage)
//         messages.GET("/search", handlers.SearchMessages)
//     }
//     // Similar for contacts, templates, etc.
// }
//
// webhooks := router.Group("/webhooks")
// {
//     webhooks.GET("/whatsapp", handlers.VerifyWebhook)
//     webhooks.POST("/whatsapp", handlers.ReceiveWebhook)
// }
//
// router.GET("/health", handlers.HealthCheck)
// router.GET("/metrics", handlers.Metrics)
```

---

## 5. Authentication & Security

### Task 5.1: API Key Authentication
**Architecture Decision:** Bearer token with bcrypt hashed keys

**Implementation Steps:**
- [ ] Create API key model (already in models)
- [ ] Implement key generation
- [ ] Implement key hashing
- [ ] Create authentication middleware
- [ ] Add key validation

**Files to Create:**
- `internal/api/middleware/auth.go`
- `pkg/utils/crypto.go`
- `internal/services/auth_service.go`

**TODO: CLAUDE_CODE**
```go
// pkg/utils/crypto.go
package utils

// TODO: CLAUDE_CODE
// Implement:
// - GenerateAPIKey() (string, error) - generates random secure key
// - HashAPIKey(key string) (string, error) - bcrypt hash
// - CompareAPIKey(hashedKey, key string) bool - verify key
// - GenerateID(prefix string) string - generate IDs like "msg_abc123"

// internal/services/auth_service.go
package services

// TODO: CLAUDE_CODE
// Implement AuthService:
// - CreateAPIKey(name string, permissions []string) (*models.APIKey, string, error)
//   Returns the model and the raw key (only time it's visible)
// - ValidateAPIKey(key string) (*models.APIKey, error)
// - RevokeAPIKey(keyID string) error
// - ListAPIKeys() ([]*models.APIKey, error)
// - UpdateAPIKey(keyID string, updates map[string]interface{}) error

// internal/api/middleware/auth.go
package middleware

// TODO: CLAUDE_CODE
// Implement:
// - AuthMiddleware() gin.HandlerFunc
//   - Extract Bearer token from Authorization header
//   - Validate format
//   - Look up API key in database
//   - Verify key hash
//   - Check expiration
//   - Update last_used_at
//   - Store key info in context
//   - Return 401 if invalid
```

---

### Task 5.2: Webhook Signature Verification
**Architecture Decision:** HMAC-SHA256 signature validation

**TODO: CLAUDE_CODE**
```go
// internal/api/middleware/webhook_auth.go
package middleware

// TODO: CLAUDE_CODE
// Implement:
// - WebhookAuthMiddleware(secret string) gin.HandlerFunc
//   - Get X-Hub-Signature-256 header
//   - Read request body
//   - Compute HMAC-SHA256
//   - Compare signatures (constant-time comparison)
//   - Store body in context for handler
//   - Return 401 if signature doesn't match

// pkg/utils/signature.go
// TODO: CLAUDE_CODE
// Implement:
// - ComputeHMAC(message, secret []byte) string
// - VerifyHMAC(message, secret []byte, signature string) bool
```

---

## 6. Service Layer

### Task 6.1: Message Service
**Architecture Decision:** Business logic in service layer

**TODO: CLAUDE_CODE**
```go
// internal/services/message_service.go
package services

// TODO: CLAUDE_CODE
// Implement MessageService struct and methods:
// - Dependencies: messageRepo, contactRepo, whatsappClient
// - SendTextMessage(phone, content string) (*models.Message, error)
// - SendMediaMessage(phone, mediaURL, caption, mediaType string) (*models.Message, error)
// - SendTemplateMessage(phone, templateName, language string, params []string) (*models.Message, error)
// - GetMessage(messageID string) (*models.Message, error)
// - ListMessages(filters MessageFilters) ([]*models.Message, *Pagination, error)
// - SearchMessages(query string, filters MessageFilters) ([]*models.Message, error)
// - ProcessIncomingMessage(webhookPayload *whatsapp.WebhookPayload) error
// - UpdateMessageStatus(whatsappMessageID, status string) error
// - GetMessagesByPhone(phone string, pagination *Pagination) ([]*models.Message, error)

// Business logic:
// - Validate phone numbers
// - Create contact if doesn't exist
// - Store message in database
// - Call WhatsApp API
// - Handle API errors
// - Update message status
```

---

### Task 6.2: Contact Service
**Architecture Decision:** Auto-create contacts, update metadata

**TODO: CLAUDE_CODE**
```go
// internal/services/contact_service.go
package services

// TODO: CLAUDE_CODE
// Implement ContactService struct and methods:
// - Dependencies: contactRepo
// - GetOrCreateContact(phone string) (*models.Contact, error)
// - GetContact(contactID string) (*models.Contact, error)
// - GetContactByPhone(phone string) (*models.Contact, error)
// - ListContacts(pagination *Pagination, filters ContactFilters) ([]*models.Contact, *Pagination, error)
// - SearchContacts(query string) ([]*models.Contact, error)
// - UpdateContact(contactID string, updates map[string]interface{}) (*models.Contact, error)
// - UpdateLastMessage(phone string, timestamp time.Time) error
// - IncrementMessageCount(phone string) error
// - UpdateUnreadCount(phone string, delta int) error
// - DeleteContact(contactID string) error

// Business logic:
// - Handle race conditions on create
// - Validate updates
// - Manage message counts atomically
```

---

### Task 6.3: Template Service
**Architecture Decision:** Manage templates with validation

**TODO: CLAUDE_CODE**
```go
// internal/services/template_service.go
package services

// TODO: CLAUDE_CODE
// Implement TemplateService struct and methods:
// - Dependencies: templateRepo
// - CreateTemplate(template *models.Template) error
// - GetTemplate(templateID string) (*models.Template, error)
// - GetTemplateByName(name, language string) (*models.Template, error)
// - ListTemplates(filters TemplateFilters) ([]*models.Template, error)
// - UpdateTemplate(templateID string, updates map[string]interface{}) error
// - DeleteTemplate(templateID string) error
// - ValidateTemplate(template *models.Template) error
// - SubstituteParameters(template *models.Template, params []string) (string, error)

// Business logic:
// - Validate template format
// - Check parameter count
// - Validate status transitions
```

---

## 7. Repository Layer

### Task 7.1: Repository Pattern Implementation
**Architecture Decision:** Repository pattern for data access

**Base Repository:**
```go
// internal/repositories/base.go
package repositories

// TODO: CLAUDE_CODE
// Implement BaseRepository with common CRUD:
// - Create(model interface{}) error
// - FindByID(id string, model interface{}) error
// - Update(id string, updates map[string]interface{}, model interface{}) error
// - Delete(id string, model interface{}) error
// - List(pagination *Pagination, model interface{}) error
```

---

### Task 7.2: Message Repository
**TODO: CLAUDE_CODE**
```go
// internal/repositories/message_repository.go
package repositories

// TODO: CLAUDE_CODE
// Implement MessageRepository:
// - Embed BaseRepository
// - FindByPhone(phone string, pagination *Pagination) ([]*models.Message, error)
// - FindByDateRange(start, end time.Time) ([]*models.Message, error)
// - FindByStatus(status string) ([]*models.Message, error)
// - FindByWhatsAppMessageID(waMessageID string) (*models.Message, error)
// - Search(query string, filters MessageFilters) ([]*models.Message, error)
// - GetStatistics() (*MessageStats, error)
// - CountByPhone(phone string) (int64, error)

// Use GORM query builder
// Optimize with indexes
// Handle pagination
```

---

### Task 7.3: Contact Repository
**TODO: CLAUDE_CODE**
```go
// internal/repositories/contact_repository.go
package repositories

// TODO: CLAUDE_CODE
// Implement ContactRepository:
// - Embed BaseRepository
// - FindByPhone(phone string) (*models.Contact, error)
// - GetOrCreate(phone string) (*models.Contact, error)
// - Search(query string) ([]*models.Contact, error)
// - UpdateLastMessage(phone string, timestamp time.Time) error
// - IncrementMessageCount(phone string, delta int) error
// - UpdateUnreadCount(phone string, delta int) error
// - FindActive(limit int) ([]*models.Contact, error)

// Use transactions where needed
// Handle race conditions with FOR UPDATE
```

---

### Task 7.4: Template Repository
**TODO: CLAUDE_CODE**
```go
// internal/repositories/template_repository.go
package repositories

// TODO: CLAUDE_CODE
// Implement TemplateRepository:
// - Embed BaseRepository
// - FindByName(name, language string) (*models.Template, error)
// - FindByCategory(category string) ([]*models.Template, error)
// - FindByStatus(status string) ([]*models.Template, error)
// - FindApproved() ([]*models.Template, error)

// Handle unique constraints
// Soft delete support
```

---

## 8. Middleware Stack

### Task 8.1: Core Middleware
**Architecture Decision:** Middleware pipeline for cross-cutting concerns

**Middleware Needed:**
1. Logging
2. Recovery (panic handling)
3. CORS
4. Request ID
5. Authentication
6. Rate Limiting
7. Request Validation
8. Error Handling

**TODO: CLAUDE_CODE**
```go
// internal/api/middleware/logging.go
package middleware

// TODO: CLAUDE_CODE
// Implement LoggingMiddleware() gin.HandlerFunc
// - Log request start
// - Log request end with duration
// - Log status code, method, path
// - Include request ID
// - Use structured logging (zap)

// internal/api/middleware/recovery.go
// TODO: CLAUDE_CODE
// Implement RecoveryMiddleware() gin.HandlerFunc
// - Recover from panics
// - Log stack trace
// - Return 500 with error ID
// - Don't expose internal errors

// internal/api/middleware/cors.go
// TODO: CLAUDE_CODE
// Implement CORSMiddleware() gin.HandlerFunc
// - Set CORS headers
// - Handle preflight OPTIONS requests
// - Configurable allowed origins

// internal/api/middleware/request_id.go
// TODO: CLAUDE_CODE
// Implement RequestIDMiddleware() gin.HandlerFunc
// - Generate unique request ID
// - Add to context
// - Add to response header
// - Use in all logs

// internal/api/middleware/rate_limit.go
// TODO: CLAUDE_CODE
// Implement RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc
// - Track requests per API key
// - Use in-memory store (or Redis future)
// - Return 429 when exceeded
// - Add rate limit headers

// internal/api/middleware/validator.go
// TODO: CLAUDE_CODE
// Implement ValidationMiddleware() gin.HandlerFunc
// - Validate request body/params
// - Use validator library
// - Return 400 with detailed errors
```

---

## 9. Configuration Management

### Task 9.1: Configuration Structure ✅ COMPLETED
**Architecture Decision:** Environment variables with Viper

**TODO: CLAUDE_CODE**
```go
// internal/config/config.go
package config

// TODO: CLAUDE_CODE
// Define Config struct with all settings:
type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    WhatsApp WhatsAppConfig
    Security SecurityConfig
    Logging  LoggingConfig
}

type ServerConfig struct {
    Port            int
    Environment     string // development, staging, production
    BaseURL         string
    ShutdownTimeout time.Duration
}

type DatabaseConfig struct {
    DSN             string
    MaxOpenConns    int
    MaxIdleConns    int
    ConnMaxLifetime time.Duration
}

type WhatsAppConfig struct {
    APIToken            string
    PhoneNumberID       string
    BusinessAccountID   string
    WebhookVerifyToken  string
    WebhookSecret       string
    APIBaseURL          string
}

type SecurityConfig struct {
    APIKeySalt     string
    SessionSecret  string
}

type LoggingConfig struct {
    Level      string // debug, info, warn, error
    Format     string // json, console
    OutputPath string
}

// TODO: CLAUDE_CODE
// Implement:
// - LoadConfig() (*Config, error)
//   - Load from environment variables
//   - Load from .env file if exists
//   - Use Viper for config management
// - Validate() error
//   - Ensure all required fields are set
//   - Validate formats
//   - Return detailed errors
// - GetDatabaseDSN() string
```

**TODO: CLAUDE_CODE**
```bash
# Create .env.example with all variables
# Document each variable
# Provide example values
# Mark required vs optional
```

---

## 10. Error Handling

### Task 10.1: Error Types and Codes ✅ COMPLETED
**Architecture Decision:** Structured errors with codes

**TODO: CLAUDE_CODE**
```go
// pkg/errors/errors.go
package errors

// TODO: CLAUDE_CODE
// Define error types:
type AppError struct {
    Code       string                 `json:"code"`
    Message    string                 `json:"message"`
    Details    map[string]interface{} `json:"details,omitempty"`
    StatusCode int                    `json:"-"`
    Err        error                  `json:"-"`
}

// Implement error constructor functions:
// - NewBadRequest(message string, details map[string]interface{}) *AppError
// - NewUnauthorized(message string) *AppError
// - NewNotFound(resource, id string) *AppError
// - NewInternalError(err error) *AppError
// - NewWhatsAppError(err error) *AppError
// - NewValidationError(errors []ValidationError) *AppError

// Error codes as constants:
const (
    ErrInvalidRequest      = "invalid_request"
    ErrUnauthorized        = "unauthorized"
    ErrNotFound            = "not_found"
    ErrInvalidPhoneNumber  = "invalid_phone_number"
    ErrWhatsAppAPI         = "whatsapp_api_error"
    ErrRateLimitExceeded   = "rate_limit_exceeded"
    // ... etc from API_DESIGN.md
)

// internal/api/middleware/error_handler.go
// TODO: CLAUDE_CODE
// Implement ErrorHandlerMiddleware() gin.HandlerFunc
// - Catch errors from handlers
// - Convert to AppError
// - Return proper JSON response
// - Log errors appropriately
// - Add request ID to response
```

---

## 11. Logging System

### Task 11.1: Structured Logging with Zap ✅ COMPLETED
**Architecture Decision:** Zap for structured JSON logging

**TODO: CLAUDE_CODE**
```go
// pkg/logger/logger.go
package logger

import "go.uber.org/zap"

// TODO: CLAUDE_CODE
// Implement:
// - InitLogger(config LoggingConfig) (*zap.Logger, error)
//   - Create logger based on environment
//   - Set log level
//   - Configure output (stdout, file)
//   - Set format (JSON for production, console for dev)
// - GetLogger() *zap.Logger
//   - Return global logger instance
// - WithFields(fields map[string]interface{}) *zap.Logger
//   - Create child logger with extra fields
// - WithRequestID(requestID string) *zap.Logger

// Usage patterns:
// logger.Info("message sent",
//     zap.String("phone", phone),
//     zap.String("message_id", msgID),
//     zap.Duration("duration", duration))
```

---

## 12. Docker & Deployment

### Task 12.1: Dockerfile
**Architecture Decision:** Multi-stage Docker build

**TODO: CLAUDE_CODE**
```dockerfile
# Dockerfile
# TODO: CLAUDE_CODE
# Create multi-stage Dockerfile:
# Stage 1: Builder
# - Use golang:1.21-alpine as base
# - Copy go.mod and go.sum
# - Download dependencies
# - Copy source code
# - Build binary with CGO_ENABLED=0
# - Strip debug symbols

# Stage 2: Runtime
# - Use alpine:latest
# - Install ca-certificates
# - Create non-root user
# - Copy binary from builder
# - Set working directory
# - Expose port 8080
# - Add healthcheck
# - Set user to non-root
# - CMD to run binary
```

---

### Task 12.2: Docker Compose
**Architecture Decision:** Docker Compose for local development

**TODO: CLAUDE_CODE**
```yaml
# docker-compose.yml
# TODO: CLAUDE_CODE
# Create docker-compose.yml with:
# Services:
#   - app:
#     - build from Dockerfile
#     - environment variables from .env
#     - depends on postgres
#     - ports 8080:8080
#     - volumes for live reload (development)
#     - restart policy
#   
#   - postgres:
#     - image: postgres:15-alpine
#     - environment variables (POSTGRES_DB, USER, PASSWORD)
#     - volumes for persistence
#     - ports 5432:5432
#     - healthcheck

# Networks:
#   - app-network (bridge)

# Volumes:
#   - postgres-data (persistent storage)
```

---

### Task 12.3: Makefile
**Architecture Decision:** Makefile for common tasks

**TODO: CLAUDE_CODE**
```makefile
# Makefile
# TODO: CLAUDE_CODE
# Create targets:
# - build: Build the binary
# - run: Run locally
# - test: Run all tests
# - test-unit: Run unit tests
# - test-integration: Run integration tests
# - docker-build: Build Docker image
# - docker-up: Start Docker Compose
# - docker-down: Stop Docker Compose
# - migrate-up: Run migrations
# - migrate-down: Rollback migrations
# - lint: Run linters
# - fmt: Format code
# - clean: Clean build artifacts
# - help: Show available targets
```

---

## Implementation Order

**Suggested Order (Most Dependencies First):**

1. **Foundation** (Can be done in parallel)
   - Project Structure (1.1, 1.2)
   - Configuration (9.1)
   - Error Handling (10.1)
   - Logging (11.1)

2. **Data Layer**
   - Database Connection (2.1)
   - Models (2.2)
   - Migrations (2.3)
   - Repositories (7.1, 7.2, 7.3, 7.4)

3. **External Integration**
   - WhatsApp Client (3.1)

4. **Business Logic**
   - Services (6.1, 6.2, 6.3)

5. **API Layer**
   - Middleware (8.1, 5.1, 5.2)
   - Handlers (covered in separate doc)
   - Routes (4.2)
   - Server Setup (4.1)

6. **Deployment**
   - Docker (12.1, 12.2, 12.3)

---

## Testing Tasks

**TODO: CLAUDE_CODE**

For each component, create corresponding test files:

```go
// Example: internal/services/message_service_test.go
package services_test

// TODO: CLAUDE_CODE
// Write tests for:
// - SendTextMessage (success, validation errors, API errors)
// - SendMediaMessage (all media types, invalid URLs)
// - ProcessIncomingMessage (various payload types)
// - UpdateMessageStatus (state transitions)
// Use testify/assert and testify/mock
// Mock repositories and WhatsApp client
```

**Test Coverage Goals:**
- Services: 80%+
- Repositories: 70%+
- Handlers: 80%+
- Middleware: 70%+

---

## Handler Implementation Tasks

**Note:** Handlers will be documented in a separate HANDLER_TASKS.md file since they map directly to API endpoints.

---

## Validation Tasks

**TODO: CLAUDE_CODE**
```go
// pkg/validator/validator.go
package validator

// TODO: CLAUDE_CODE
// Implement validators:
// - ValidatePhoneNumber(phone string) error
//   - E.164 format check
//   - Length validation
// - ValidateURL(url string) error
// - ValidateMessageType(msgType string) error
// - ValidateStatus(status string) error
// - ValidateLanguageCode(code string) error
// - ValidateEmail(email string) error
```

---

## Utility Functions

**TODO: CLAUDE_CODE**
```go
// pkg/utils/pagination.go
package utils

// TODO: CLAUDE_CODE
// Implement:
// - Pagination struct (Limit, Offset, Total, HasMore)
// - NewPagination(limit, offset int) *Pagination
// - ApplyToQuery(db *gorm.DB) *gorm.DB
// - SetTotal(total int64)
// - ToResponse() PaginationResponse

// pkg/utils/response.go
// TODO: CLAUDE_CODE
// Implement response helpers:
// - SuccessResponse(c *gin.Context, status int, data interface{})
// - ErrorResponse(c *gin.Context, err *errors.AppError)
// - ListResponse(c *gin.Context, data interface{}, pagination *Pagination)
```

---

## Notes

- All TODO: CLAUDE_CODE markers indicate where actual implementation is needed
- Follow Go best practices and idioms
- Use dependency injection for testability
- Keep business logic in services, not handlers
- Use interfaces for external dependencies
- Write tests alongside implementation
- Document exported functions and types
- Use meaningful variable names
- Handle errors explicitly
- Don't panic in production code (except in init functions)

---

**Status:** Ready for implementation  
**Next:** Create HANDLER_TASKS.md for API endpoint handlers
