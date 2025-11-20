# Core Features
# Vibecoded WA Client

**Last Updated:** November 21, 2025

---

## Feature Categories

### 1. Authentication & Configuration
- WhatsApp Business API authentication
- Environment-based configuration
- API key management for REST endpoints
- Phone number registration and verification
- Configuration validation on startup

**TODO for Claude Code:** Implementation of config validation module

---

### 2. User Management

#### User Registration & Authentication
- User account creation
- Password hashing and validation
- JWT token generation and validation
- Password reset functionality
- Email verification (optional)

#### User Roles & Permissions
- Admin role (full access)
- User role (limited access)
- API-only role (no dashboard access)
- Custom permission sets

#### User Management Operations
- Create new users
- Update user profiles
- Deactivate/reactivate users
- Delete users
- List users with filtering

**TODO for Claude Code:** User authentication middleware
**TODO for Claude Code:** Role-based access control implementation

---

### 3. API Key Management

#### Features
- Generate API keys for programmatic access
- Multiple keys per user
- Key naming and description
- Key expiration dates
- Rate limiting per key
- Key usage tracking
- Key revocation

#### Security
- Keys stored as hashed values
- Bearer token authentication
- Key rotation support
- Usage audit logs

**TODO for Claude Code:** API key generation and validation service

---

### 4. Send Messages

#### Text Messages
- Send plain text to phone numbers
- Support international phone format
- Handle message length limits
- Return message ID

#### Media Messages
- Images (JPEG, PNG)
- Documents (PDF, DOCX, etc.)
- Support for media URLs
- File size validation

#### Template Messages
- Send pre-approved templates
- Parameter substitution
- Template management

**TODO for Claude Code:** Message sending service with retry logic

---

### 5. Receive Messages (Webhooks)

#### Webhook Handler
- Webhook endpoint for WhatsApp events
- Signature verification
- Message parsing and validation
- Idempotent message processing

#### Supported Events
- Incoming messages
- Message status updates
- Delivery receipts
- Read receipts

**TODO for Claude Code:** Webhook signature verification
**TODO for Claude Code:** Idempotency key handling

---

### 6. Contact Management

#### Features
- Store contact information
- Automatic contact creation from messages
- Update contact profiles
- Search contacts by phone/name
- Contact metadata storage
- Contact tags and categories
- Last interaction tracking

**TODO for Claude Code:** Contact auto-creation from messages

---

### 7. Message History

#### Storage
- Persist all sent/received messages
- Message status tracking
- Error logging for failed messages

#### Querying
- Query message history
- Pagination support
- Filter by phone number, date range, message type, status
- Full-text search capability
- Sort by timestamp

**TODO for Claude Code:** Optimized message query builder
**TODO for Claude Code:** Full-text search implementation

---

### 8. Core Infrastructure

#### Database
- PostgreSQL integration
- Database migrations
- Connection pooling
- Query optimization

#### Logging & Monitoring
- Structured logging (JSON format)
- Request/response logging
- Error tracking
- Health check endpoint

#### Error Handling
- Graceful error recovery
- Retry logic with exponential backoff
- Circuit breaker pattern
- Graceful shutdown

**TODO for Claude Code:** Database migration system
**TODO for Claude Code:** Structured logging middleware

---

### 9. Template Management
- CRUD operations for templates
- Template validation
- Parameter management
- Template categories
- Template approval workflow

**TODO for Claude Code:** Template validation service

---

### 10. Advanced Querying
- Message analytics
- Conversation threads
- Unread message tracking
- Contact activity tracking
- Delivery rate statistics

**TODO for Claude Code:** Analytics aggregation queries

---

### 11. Rate Limiting & Queue

#### Rate Limiting
- Request rate limiting per API key
- Per-user rate limits
- Rate limit headers in responses
- Rate limit configuration

#### Message Queue
- Message queue for bulk sending
- Retry logic for failed messages
- Priority queue support
- Dead letter queue

**TODO for Claude Code:** Rate limiting middleware
**TODO for Claude Code:** Message queue implementation

---

### 12. MCP Server Core

#### Implementation
- JSON-RPC 2.0 implementation
- Tool registration system
- stdio transport support

#### Basic Tools
- send_whatsapp_message
- get_whatsapp_messages
- list_whatsapp_contacts
- search_messages
- get_user_info

**TODO for Claude Code:** MCP JSON-RPC server
**TODO for Claude Code:** MCP tool registry system

---

### 13. Observability

#### Metrics
- Prometheus metrics endpoint
- Request/response metrics
- Error rate tracking
- Database query performance
- Message delivery rates

#### Monitoring
- Health check with detailed status
- Resource usage tracking
- Alert thresholds

**TODO for Claude Code:** Prometheus metrics integration

---

### 14. Interactive Messages
- Button messages
- List messages
- Reply buttons
- Quick replies

**TODO for Claude Code:** Interactive message builder

---

### 15. Advanced Media
- Voice messages
- Location sharing
- Contact cards
- Video messages
- Sticker support

**TODO for Claude Code:** Media type handlers

---

### 16. Business Profile
- Update business profile
- Business hours
- About information
- Profile picture management

**TODO for Claude Code:** Business profile API integration

---

### 17. Advanced MCP Tools
- Conversation management
- Template creation via Claude
- Analytics queries
- Bulk operations
- User management via Claude

**TODO for Claude Code:** Additional MCP tools

---

### 18. Admin Features
- Configuration API
- System settings management
- Webhook management
- Backup and restore
- System diagnostics

**TODO for Claude Code:** Admin API endpoints

---

## Feature Dependencies

```
User Management → API Keys → Rate Limiting
                          ↓
Send Messages → Message History → Search
             ↓
Contact Management → Contact Search
             ↓
Webhooks → Message Receiving → Contact Auto-creation
                             ↓
                      Message Threading
```

---

## Testing Requirements

Each feature requires:
- ✅ Unit tests for business logic
- ✅ Integration tests for API endpoints
- ✅ End-to-end tests for critical flows
- ✅ Performance tests for high-load scenarios
- ✅ Security tests for authentication/authorization

**TODO for Claude Code:** Test suite for each feature module
