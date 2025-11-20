# Implementation Tasks
# Vibecoded WA Client

**Last Updated:** November 21, 2025

---

## Overview

This document provides a complete breakdown of implementation tasks. Each section includes clear TODO markers for Claude Code.

---

## 1. Foundation ✅ COMPLETED

### Project Structure
- [x] Directory structure
- [x] Go module initialization
- [x] Package organization

### Configuration
- [x] Config structure with Viper
- [x] Environment variable loading
- [x] Validation system
- [ ] Feature flags (optional enhancement)

### Logging
- [x] Zap structured logging
- [x] Multiple output formats
- [x] Request ID tracking

### Error Handling
- [x] Custom error types
- [x] HTTP status mapping
- [x] Error response formatting

### Utilities
- [x] Pagination helpers
- [x] Response helpers
- [x] Crypto utilities (API key generation, hashing)
- [x] Input validators

---

## 2. Database Layer ✅ PARTIALLY COMPLETED

### Connection Management ✅
- [x] Database connection with GORM
- [x] Connection pooling
- [x] Health checks
- [x] Auto-migration support

### Data Models ✅
- [x] Message model
- [x] Contact model
- [x] Template model
- [x] APIKey model
- [x] Call & Transcript models (future use)
- [x] Custom JSONB types

### Migrations
**TODO: CLAUDE_CODE**
- [ ] Create migration files for all tables
- [ ] Add indexes as specified in DATA_MODELS.md
- [ ] Create triggers for updated_at timestamps
- [ ] Migration runner script
- [ ] Rollback support

### Repository Layer
**TODO: CLAUDE_CODE**

#### Base Repository
- [ ] BaseRepository with common CRUD operations
- [ ] Generic query builder
- [ ] Transaction support helpers

#### Message Repository
- [ ] FindByPhone with pagination
- [ ] FindByDateRange
- [ ] FindByStatus
- [ ] FindByWhatsAppMessageID
- [ ] Full-text search implementation
- [ ] GetStatistics aggregations
- [ ] CountByPhone

#### Contact Repository
- [ ] FindByPhone
- [ ] GetOrCreate with race condition handling
- [ ] Search functionality
- [ ] UpdateLastMessage atomic operation
- [ ] IncrementMessageCount atomic operation
- [ ] UpdateUnreadCount atomic operation
- [ ] FindActive contacts

#### Template Repository
- [ ] FindByName with language
- [ ] FindByCategory
- [ ] FindByStatus
- [ ] FindApproved templates
- [ ] Soft delete support

---

## 3. WhatsApp Integration

### WhatsApp Client
**TODO: CLAUDE_CODE**

#### Client Setup
- [ ] WhatsAppClient struct with HTTP client
- [ ] Authentication configuration
- [ ] Timeout and retry logic
- [ ] Error response parsing

#### Message Operations
- [ ] SendTextMessage
- [ ] SendMediaMessage (image, document, audio, video)
- [ ] SendTemplateMessage with parameter substitution
- [ ] GetMessageStatus

#### Error Handling
- [ ] Parse WhatsApp API error responses
- [ ] Map to application error types
- [ ] Retry logic for transient failures
- [ ] Circuit breaker pattern

### Webhook Processing
**TODO: CLAUDE_CODE**
- [ ] HMAC-SHA256 signature verification
- [ ] Webhook payload parsing
- [ ] Message event extraction
- [ ] Status update event extraction
- [ ] Call event extraction (future)
- [ ] Idempotency handling

---

## 4. Core Services

### Message Service
**TODO: CLAUDE_CODE**

#### Sending Operations
- [ ] SendTextMessage
  - Phone number validation
  - Content length validation
  - WhatsApp API call
  - Database storage
  - Auto-create contact if needed
  
- [ ] SendMediaMessage
  - URL validation and accessibility check
  - Media type validation
  - File size limits enforcement
  - WhatsApp API call
  - Store media metadata
  
- [ ] SendTemplateMessage
  - Template lookup and validation
  - Parameter count validation
  - Parameter substitution
  - WhatsApp API call
  - Database storage

#### Query Operations
- [ ] GetMessage by ID
- [ ] ListMessages with filters and pagination
- [ ] SearchMessages with full-text search
- [ ] GetMessagesByPhone with pagination

#### Webhook Processing
- [ ] ProcessIncomingMessage
  - Parse webhook payload
  - Get or create contact
  - Store message in database
  - Update contact's last_message_at
  - Update unread count
  
- [ ] UpdateMessageStatus
  - Find message by WhatsApp message ID
  - Update status field
  - Update timestamp

### Contact Service
**TODO: CLAUDE_CODE**
- [ ] GetOrCreateContact with race condition handling
- [ ] GetContact by ID
- [ ] GetContactByPhone
- [ ] ListContacts with pagination and sorting
- [ ] SearchContacts by name or phone
- [ ] UpdateContact with validation
- [ ] UpdateLastMessage timestamp
- [ ] IncrementMessageCount atomically
- [ ] UpdateUnreadCount atomically

### Template Service
**TODO: CLAUDE_CODE**
- [ ] CreateTemplate with validation
- [ ] GetTemplate by ID
- [ ] GetTemplateByName with language
- [ ] ListTemplates with filters
- [ ] UpdateTemplate
- [ ] DeleteTemplate (soft delete)
- [ ] ValidateTemplate structure
- [ ] SubstituteParameters in template content

### Auth Service
**TODO: CLAUDE_CODE**
- [ ] CreateAPIKey (returns raw key once)
- [ ] ValidateAPIKey against hashed version
- [ ] RevokeAPIKey
- [ ] ListAPIKeys
- [ ] UpdateAPIKey metadata
- [ ] TrackAPIKeyUsage

---

## 5. API Handlers

See [HANDLER_TASKS.md](HANDLER_TASKS.md) for detailed handler implementation tasks.

### Messages Handlers
**TODO: CLAUDE_CODE**
- [ ] SendMessage - POST /api/v1/messages
- [ ] GetMessage - GET /api/v1/messages/:id
- [ ] ListMessages - GET /api/v1/messages
- [ ] SearchMessages - GET /api/v1/messages/search

### Contacts Handlers
**TODO: CLAUDE_CODE**
- [ ] ListContacts - GET /api/v1/contacts
- [ ] GetContact - GET /api/v1/contacts/:id
- [ ] UpdateContact - PATCH /api/v1/contacts/:id
- [ ] SearchContacts - GET /api/v1/contacts/search

### Templates Handlers
**TODO: CLAUDE_CODE**
- [ ] ListTemplates - GET /api/v1/templates
- [ ] GetTemplate - GET /api/v1/templates/:id
- [ ] CreateTemplate - POST /api/v1/templates
- [ ] UpdateTemplate - PATCH /api/v1/templates/:id
- [ ] DeleteTemplate - DELETE /api/v1/templates/:id

### Webhook Handlers
**TODO: CLAUDE_CODE**
- [ ] VerifyWebhook - GET /webhooks/whatsapp
- [ ] ReceiveWebhook - POST /webhooks/whatsapp

### System Handlers
**TODO: CLAUDE_CODE**
- [ ] HealthCheck - GET /health
- [ ] Metrics - GET /metrics (Prometheus)

### Call Handlers (Future)
- [ ] Create stubs for all call endpoints
- [ ] Return 501 Not Implemented

---

## 6. Middleware

### Authentication & Security
**TODO: CLAUDE_CODE**
- [ ] AuthMiddleware for API key validation
- [ ] WebhookAuthMiddleware for signature verification
- [ ] CORS middleware
- [ ] Security headers middleware

### Request Processing
**TODO: CLAUDE_CODE**
- [ ] LoggingMiddleware with structured logs
- [ ] RequestIDMiddleware
- [ ] RecoveryMiddleware for panic handling
- [ ] ValidationMiddleware for request validation
- [ ] RateLimitMiddleware per API key
- [ ] TimeoutMiddleware for long requests
- [ ] CompressionMiddleware for responses

### Error Handling
**TODO: CLAUDE_CODE**
- [ ] ErrorHandlerMiddleware
- [ ] Error to HTTP status mapping
- [ ] Detailed error responses with codes

---

## 7. API Server Setup

**TODO: CLAUDE_CODE**

### Server Initialization
- [ ] Server struct with dependencies
- [ ] Gin engine initialization
- [ ] Middleware registration
- [ ] Route registration
- [ ] Graceful shutdown handling
- [ ] Signal handling (SIGTERM, SIGINT)

### Route Registration
- [ ] Create route groups (/api/v1/messages, /api/v1/contacts, etc.)
- [ ] Apply middleware to appropriate groups
- [ ] Register all handler functions
- [ ] Document route structure

### Main Entry Point
- [ ] Load configuration
- [ ] Initialize database connection
- [ ] Initialize WhatsApp client
- [ ] Create all services
- [ ] Create all handlers
- [ ] Start server
- [ ] Handle shutdown cleanup

---

## 8. Testing

### Unit Tests
**TODO: CLAUDE_CODE**
- [ ] Service layer tests (80%+ coverage)
- [ ] Repository layer tests (70%+ coverage)
- [ ] Utility function tests
- [ ] Middleware tests

### Integration Tests
**TODO: CLAUDE_CODE**
- [ ] API endpoint tests
- [ ] Database integration tests
- [ ] WhatsApp client mock tests
- [ ] Webhook processing tests

### End-to-End Tests
**TODO: CLAUDE_CODE**
- [ ] Full message send/receive flow
- [ ] Contact auto-creation flow
- [ ] Template message flow
- [ ] Error handling scenarios

### Test Infrastructure
**TODO: CLAUDE_CODE**
- [ ] Test database setup/teardown
- [ ] Test fixtures and factories
- [ ] Mock implementations
- [ ] Test helpers and utilities
- [ ] Coverage reporting

---

## 9. Deployment

### Docker Setup
**TODO: CLAUDE_CODE**
- [ ] Multi-stage Dockerfile
  - Builder stage with Go compilation
  - Runtime stage with minimal image
  - Non-root user configuration
  - Health check inclusion
  
- [ ] Docker Compose configuration
  - Application service
  - PostgreSQL service
  - Network configuration
  - Volume mounts
  - Environment variables

### Build Scripts
**TODO: CLAUDE_CODE**
- [ ] Makefile with common tasks
  - build: Compile binary
  - run: Run locally
  - test: Run all tests
  - docker-build: Build Docker image
  - docker-up: Start containers
  - docker-down: Stop containers
  - migrate-up: Run migrations
  - migrate-down: Rollback
  - lint: Run linters
  - fmt: Format code
  - clean: Clean artifacts

### Deployment Documentation
**TODO: CLAUDE_CODE**
- [ ] Local development setup guide
- [ ] Docker deployment guide
- [ ] Environment variable documentation
- [ ] Database migration guide
- [ ] Troubleshooting guide

---

## 10. Documentation

### API Documentation
**TODO: CLAUDE_CODE**
- [ ] OpenAPI/Swagger specification
- [ ] API endpoint documentation with examples
- [ ] Error code documentation
- [ ] Authentication documentation
- [ ] Rate limiting documentation

### Code Documentation
**TODO: CLAUDE_CODE**
- [ ] Godoc comments for all exported functions
- [ ] Package documentation
- [ ] Example code snippets
- [ ] Architecture decision records

### User Documentation
**TODO: CLAUDE_CODE**
- [ ] README with quick start
- [ ] Installation guide
- [ ] Configuration guide
- [ ] API usage examples
- [ ] Troubleshooting guide
- [ ] FAQ

---

## 11. Future Enhancements

### Calling Features
- [ ] Call initiation
- [ ] Call management (mute, hold, transfer)
- [ ] Call recording
- [ ] Call transcription (Deepgram integration)
- [ ] Text-to-speech (11labs integration)
- [ ] WebRTC signaling

### MCP Server
- [ ] JSON-RPC 2.0 server implementation
- [ ] Tool registration system
- [ ] MCP tools for messaging
- [ ] MCP tools for contacts
- [ ] MCP tools for analytics

### Web UI
- [ ] React/Vue dashboard
- [ ] Message history view
- [ ] Contact management UI
- [ ] Analytics dashboard
- [ ] Configuration UI

### Advanced Features
- [ ] Interactive messages (buttons, lists)
- [ ] Message templates management UI
- [ ] Bulk message sending
- [ ] Message scheduling
- [ ] Analytics and reporting
- [ ] Webhook forwarding
- [ ] Multi-account support

---

## Implementation Notes

- **Start with Foundation**: Ensure all foundation components are solid before building features
- **Test as You Go**: Write tests alongside implementation
- **Follow Patterns**: Use consistent patterns across all components
- **Document Decisions**: Update architecture docs when making significant decisions
- **Iterate**: Build MVP first, then enhance

---

## Getting Started

1. Review architecture decisions in [ARCHITECTURE.md](ARCHITECTURE.md)
2. Understand data models in [DATA_MODELS.md](DATA_MODELS.md)
3. Check API design in [API_DESIGN.md](API_DESIGN.md)
4. Start with database migrations
5. Implement repositories
6. Build services
7. Create handlers
8. Add middleware
9. Wire everything together
10. Test thoroughly

---

**Status:** Ready for implementation  
**Maintained by:** Ashok  
**Vibecoded:** Yes 🎵
