# API Handler Implementation Tasks
# Vibecoded WA Client

**Last Updated:** November 20, 2025  
**Status:** 🟡 Planning Phase

---

## Overview

This document breaks down all API endpoint handlers that need to be implemented. Each handler corresponds to an endpoint defined in API_DESIGN.md.

**Total Endpoints:** 27

---

## Table of Contents

1. [Messages Handlers](#1-messages-handlers)
2. [Contacts Handlers](#2-contacts-handlers)
3. [Templates Handlers](#3-templates-handlers)
4. [Calls Handlers](#4-calls-handlers)
5. [Webhook Handlers](#5-webhook-handlers)
6. [System Handlers](#6-system-handlers)
7. [Request/Response Models](#7-requestresponse-models)

---

## Handler Structure

All handlers follow this pattern:

```go
// internal/api/handlers/handler_name.go
package handlers

type HandlerName struct {
    service ServiceInterface
    logger  *zap.Logger
}

func NewHandlerName(service ServiceInterface, logger *zap.Logger) *HandlerName {
    return &HandlerName{
        service: service,
        logger:  logger,
    }
}

func (h *HandlerName) MethodName(c *gin.Context) {
    // 1. Parse and validate request
    // 2. Call service layer
    // 3. Handle errors
    // 4. Return response
}
```

---

## 1. Messages Handlers

### File: `internal/api/handlers/messages.go`

### Handler 1.1: Send Message
**Endpoint:** `POST /api/v1/messages`  
**Description:** Send text, media, or template message

**TODO: CLAUDE_CODE**
```go
// internal/api/handlers/messages.go
package handlers

type MessagesHandler struct {
    messageService services.MessageService
    logger         *zap.Logger
}

func NewMessagesHandler(messageService services.MessageService, logger *zap.Logger) *MessagesHandler {
    return &MessagesHandler{
        messageService: messageService,
        logger:         logger,
    }
}

// TODO: CLAUDE_CODE
// Implement SendMessage handler
func (h *MessagesHandler) SendMessage(c *gin.Context) {
    // 1. Parse request body into SendMessageRequest
    // 2. Validate request:
    //    - Required fields based on message type
    //    - Phone number format
    //    - Content length limits
    //    - Media URL if applicable
    // 3. Extract API key from context (set by auth middleware)
    // 4. Call appropriate service method based on type:
    //    - text: messageService.SendTextMessage()
    //    - image/document/audio/video: messageService.SendMediaMessage()
    //    - template: messageService.SendTemplateMessage()
    // 5. Handle errors:
    //    - Validation errors → 400
    //    - WhatsApp API errors → 502
    //    - Internal errors → 500
    // 6. Log the operation
    // 7. Return 201 with message details
}
```

**Request Models:**
```go
// TODO: CLAUDE_CODE
type SendMessageRequest struct {
    Phone            string   `json:"phone" binding:"required"`
    Type             string   `json:"type" binding:"required,oneof=text image document audio video template"`
    Content          string   `json:"content,omitempty"`
    MediaURL         string   `json:"media_url,omitempty"`
    Caption          string   `json:"caption,omitempty"`
    Filename         string   `json:"filename,omitempty"`
    TemplateName     string   `json:"template_name,omitempty"`
    TemplateLanguage string   `json:"template_language,omitempty"`
    Parameters       []string `json:"parameters,omitempty"`
}
```

---

### Handler 1.2: Get Message by ID
**Endpoint:** `GET /api/v1/messages/:id`  
**Description:** Retrieve specific message

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *MessagesHandler) GetMessage(c *gin.Context) {
    // 1. Extract message ID from URL params
    // 2. Validate ID format
    // 3. Call messageService.GetMessage(messageID)
    // 4. Handle errors:
    //    - Not found → 404
    //    - Internal error → 500
    // 5. Return 200 with message details
}
```

---

### Handler 1.3: List Messages
**Endpoint:** `GET /api/v1/messages`  
**Description:** List messages with filters and pagination

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *MessagesHandler) ListMessages(c *gin.Context) {
    // 1. Parse query parameters:
    //    - phone (optional)
    //    - direction (optional)
    //    - type (optional)
    //    - status (optional)
    //    - start_date (optional, parse ISO 8601)
    //    - end_date (optional, parse ISO 8601)
    //    - limit (default: 20, max: 100)
    //    - offset (default: 0)
    // 2. Validate parameters:
    //    - Valid date formats
    //    - Valid enum values
    //    - Reasonable limit/offset
    // 3. Create filters struct
    // 4. Call messageService.ListMessages(filters)
    // 5. Get total count for pagination
    // 6. Return 200 with:
    //    - data: array of messages
    //    - pagination: { total, limit, offset, has_more }
}
```

**Query Models:**
```go
// TODO: CLAUDE_CODE
type ListMessagesQuery struct {
    Phone     string    `form:"phone"`
    Direction string    `form:"direction" binding:"omitempty,oneof=inbound outbound"`
    Type      string    `form:"type"`
    Status    string    `form:"status"`
    StartDate time.Time `form:"start_date" time_format:"2006-01-02"`
    EndDate   time.Time `form:"end_date" time_format:"2006-01-02"`
    Limit     int       `form:"limit" binding:"omitempty,min=1,max=100"`
    Offset    int       `form:"offset" binding:"omitempty,min=0"`
}
```

---

### Handler 1.4: Search Messages
**Endpoint:** `GET /api/v1/messages/search`  
**Description:** Full-text search in messages

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *MessagesHandler) SearchMessages(c *gin.Context) {
    // 1. Parse query parameters:
    //    - q (required, search query)
    //    - phone (optional, limit to specific contact)
    //    - limit (default: 20)
    // 2. Validate:
    //    - Query not empty
    //    - Reasonable limit
    // 3. Call messageService.SearchMessages(query, filters)
    // 4. Return 200 with:
    //    - data: array of matching messages
    //    - count: total matches
}
```

---

## 2. Contacts Handlers

### File: `internal/api/handlers/contacts.go`

### Handler 2.1: List Contacts
**Endpoint:** `GET /api/v1/contacts`  
**Description:** List all contacts with pagination

**TODO: CLAUDE_CODE**
```go
// internal/api/handlers/contacts.go
package handlers

type ContactsHandler struct {
    contactService services.ContactService
    logger         *zap.Logger
}

func NewContactsHandler(contactService services.ContactService, logger *zap.Logger) *ContactsHandler {
    return &ContactsHandler{
        contactService: contactService,
        logger:         logger,
    }
}

// TODO: CLAUDE_CODE
func (h *ContactsHandler) ListContacts(c *gin.Context) {
    // 1. Parse query parameters:
    //    - limit (default: 50, max: 100)
    //    - offset (default: 0)
    //    - sort (default: "last_message_at")
    //    - order (default: "desc", options: "asc", "desc")
    // 2. Validate parameters
    // 3. Call contactService.ListContacts(filters, pagination)
    // 4. Return 200 with:
    //    - data: array of contacts
    //    - pagination: { total, limit, offset, has_more }
}
```

---

### Handler 2.2: Get Contact
**Endpoint:** `GET /api/v1/contacts/:id`  
**Description:** Get contact details

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *ContactsHandler) GetContact(c *gin.Context) {
    // 1. Extract contact ID from URL params
    // 2. Validate ID format
    // 3. Call contactService.GetContact(contactID)
    // 4. Handle errors:
    //    - Not found → 404
    //    - Internal error → 500
    // 5. Return 200 with contact details
}
```

---

### Handler 2.3: Update Contact
**Endpoint:** `PATCH /api/v1/contacts/:id`  
**Description:** Update contact information

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *ContactsHandler) UpdateContact(c *gin.Context) {
    // 1. Extract contact ID from URL params
    // 2. Parse request body into UpdateContactRequest
    // 3. Validate:
    //    - At least one field to update
    //    - Valid metadata structure if provided
    // 4. Call contactService.UpdateContact(contactID, updates)
    // 5. Handle errors:
    //    - Not found → 404
    //    - Validation errors → 400
    //    - Internal error → 500
    // 6. Return 200 with updated contact
}
```

**Request Model:**
```go
// TODO: CLAUDE_CODE
type UpdateContactRequest struct {
    Name     *string                `json:"name,omitempty"`
    Metadata map[string]interface{} `json:"metadata,omitempty"`
}
```

---

### Handler 2.4: Search Contacts
**Endpoint:** `GET /api/v1/contacts/search`  
**Description:** Search contacts by name or phone

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *ContactsHandler) SearchContacts(c *gin.Context) {
    // 1. Parse query parameters:
    //    - q (required, search query)
    //    - limit (optional, default: 50)
    // 2. Validate query not empty
    // 3. Call contactService.SearchContacts(query)
    // 4. Return 200 with:
    //    - data: array of matching contacts
    //    - count: total matches
}
```

---

## 3. Templates Handlers

### File: `internal/api/handlers/templates.go`

### Handler 3.1: List Templates
**Endpoint:** `GET /api/v1/templates`  
**Description:** List all templates

**TODO: CLAUDE_CODE**
```go
// internal/api/handlers/templates.go
package handlers

type TemplatesHandler struct {
    templateService services.TemplateService
    logger          *zap.Logger
}

func NewTemplatesHandler(templateService services.TemplateService, logger *zap.Logger) *TemplatesHandler {
    return &TemplatesHandler{
        templateService: templateService,
        logger:          logger,
    }
}

// TODO: CLAUDE_CODE
func (h *TemplatesHandler) ListTemplates(c *gin.Context) {
    // 1. Parse optional filters:
    //    - category
    //    - status
    //    - language
    // 2. Call templateService.ListTemplates(filters)
    // 3. Return 200 with array of templates
}
```

---

### Handler 3.2: Get Template
**Endpoint:** `GET /api/v1/templates/:id`  
**Description:** Get template details

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *TemplatesHandler) GetTemplate(c *gin.Context) {
    // 1. Extract template ID from URL params
    // 2. Call templateService.GetTemplate(templateID)
    // 3. Handle not found → 404
    // 4. Return 200 with template details
}
```

---

### Handler 3.3: Create Template
**Endpoint:** `POST /api/v1/templates`  
**Description:** Create new template

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *TemplatesHandler) CreateTemplate(c *gin.Context) {
    // 1. Parse request body into CreateTemplateRequest
    // 2. Validate:
    //    - Required fields (name, language, category, content)
    //    - Valid category/status values
    //    - Parameter array format
    // 3. Call templateService.CreateTemplate(template)
    // 4. Handle errors:
    //    - Duplicate name → 400
    //    - Validation errors → 400
    //    - Internal error → 500
    // 5. Return 201 with created template
}
```

**Request Model:**
```go
// TODO: CLAUDE_CODE
type CreateTemplateRequest struct {
    Name       string   `json:"name" binding:"required"`
    Language   string   `json:"language" binding:"required"`
    Category   string   `json:"category" binding:"required,oneof=utility marketing authentication"`
    Content    string   `json:"content" binding:"required"`
    Parameters []string `json:"parameters,omitempty"`
}
```

---

### Handler 3.4: Update Template
**Endpoint:** `PATCH /api/v1/templates/:id`  
**Description:** Update template

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *TemplatesHandler) UpdateTemplate(c *gin.Context) {
    // 1. Extract template ID from URL params
    // 2. Parse request body into UpdateTemplateRequest
    // 3. Validate updates
    // 4. Call templateService.UpdateTemplate(templateID, updates)
    // 5. Handle errors (not found, validation, etc.)
    // 6. Return 200 with updated template
}
```

---

### Handler 3.5: Delete Template
**Endpoint:** `DELETE /api/v1/templates/:id`  
**Description:** Delete template (soft delete)

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *TemplatesHandler) DeleteTemplate(c *gin.Context) {
    // 1. Extract template ID from URL params
    // 2. Call templateService.DeleteTemplate(templateID)
    // 3. Handle not found → 404
    // 4. Return 204 No Content
}
```

---

## 4. Calls Handlers

### File: `internal/api/handlers/calls.go`

**Note:** Call handlers are for future implementation. Create stubs for now.

### Handler 4.1: Initiate Call
**Endpoint:** `POST /api/v1/calls`  
**Description:** Start outbound call

**TODO: CLAUDE_CODE**
```go
// internal/api/handlers/calls.go
package handlers

// TODO: CLAUDE_CODE - Future implementation
// This will require:
// - WhatsApp calling API integration
// - WebRTC signaling
// - Call state management

func (h *CallsHandler) InitiateCall(c *gin.Context) {
    // Stub for future implementation
    c.JSON(501, gin.H{"error": "Calling feature not yet implemented"})
}
```

---

### Handler 4.2-4.10: Other Call Handlers
**TODO: CLAUDE_CODE**
```go
// Create stubs for all call endpoints:
// - ListCalls
// - GetCall
// - EndCall
// - StartRecording
// - StopRecording
// - GetRecording
// - GetTranscription
// - SendTTS
// - ControlCall

// All return 501 Not Implemented for now
```

---

## 5. Webhook Handlers

### File: `internal/api/handlers/webhooks.go`

### Handler 5.1: Verify Webhook
**Endpoint:** `GET /webhooks/whatsapp`  
**Description:** WhatsApp webhook verification

**TODO: CLAUDE_CODE**
```go
// internal/api/handlers/webhooks.go
package handlers

type WebhooksHandler struct {
    messageService services.MessageService
    logger         *zap.Logger
    verifyToken    string // from config
}

func NewWebhooksHandler(messageService services.MessageService, verifyToken string, logger *zap.Logger) *WebhooksHandler {
    return &WebhooksHandler{
        messageService: messageService,
        verifyToken:    verifyToken,
        logger:         logger,
    }
}

// TODO: CLAUDE_CODE
func (h *WebhooksHandler) VerifyWebhook(c *gin.Context) {
    // 1. Parse query parameters:
    //    - hub.mode (should be "subscribe")
    //    - hub.challenge (random string)
    //    - hub.verify_token (matches our verify token)
    // 2. Validate:
    //    - mode is "subscribe"
    //    - verify_token matches configured token
    // 3. If valid: return hub.challenge as plain text (200)
    // 4. If invalid: return 403 Forbidden
}
```

---

### Handler 5.2: Receive Webhook
**Endpoint:** `POST /webhooks/whatsapp`  
**Description:** Receive WhatsApp events

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *WebhooksHandler) ReceiveWebhook(c *gin.Context) {
    // Note: Signature verification is done in middleware
    
    // 1. Get parsed webhook body from context (set by middleware)
    // 2. Identify event type:
    //    - Message received
    //    - Message status update
    //    - Call event (future)
    // 3. Route to appropriate handler:
    //    - For messages: call messageService.ProcessIncomingMessage()
    //    - For status: call messageService.UpdateMessageStatus()
    // 4. Handle idempotency (check if already processed)
    // 5. Log the event
    // 6. Return 200 OK immediately (WhatsApp expects fast response)
    // 7. Process in background if needed
}
```

**Helper Methods:**
```go
// TODO: CLAUDE_CODE
func (h *WebhooksHandler) handleMessageEvent(event *whatsapp.MessageEvent) error {
    // Process incoming message
    // - Extract message details
    // - Create/update contact
    // - Store message
    // - Trigger any automation
}

func (h *WebhooksHandler) handleStatusEvent(event *whatsapp.StatusEvent) error {
    // Update message status
    // - Find message by WhatsApp message ID
    // - Update status
    // - Update timestamp
}

func (h *WebhooksHandler) isAlreadyProcessed(messageID string) bool {
    // Check if we've already processed this webhook
    // Prevent duplicate message creation
}
```

---

## 6. System Handlers

### File: `internal/api/handlers/system.go`

### Handler 6.1: Health Check
**Endpoint:** `GET /health`  
**Description:** System health status

**TODO: CLAUDE_CODE**
```go
// internal/api/handlers/system.go
package handlers

type SystemHandler struct {
    db             *gorm.DB
    whatsappClient *whatsapp.WhatsAppClient
    version        string
    startTime      time.Time
    logger         *zap.Logger
}

func NewSystemHandler(db *gorm.DB, whatsappClient *whatsapp.WhatsAppClient, version string, logger *zap.Logger) *SystemHandler {
    return &SystemHandler{
        db:             db,
        whatsappClient: whatsappClient,
        version:        version,
        startTime:      time.Now(),
        logger:         logger,
    }
}

// TODO: CLAUDE_CODE
func (h *SystemHandler) HealthCheck(c *gin.Context) {
    // 1. Check database connection:
    //    - db.DB().Ping()
    // 2. Check WhatsApp API (optional):
    //    - Simple API call to verify connectivity
    // 3. Calculate uptime:
    //    - time.Since(h.startTime)
    // 4. Return status:
    //    - 200 if all healthy
    //    - 503 if any check fails
    // 5. Response format:
    //    {
    //      "status": "healthy" or "unhealthy",
    //      "version": "0.1.0",
    //      "uptime": "24h30m15s",
    //      "checks": {
    //        "database": "connected" or "error: ...",
    //        "whatsapp_api": "reachable" or "error: ..."
    //      },
    //      "timestamp": "2025-11-20T10:30:00Z"
    //    }
}
```

---

### Handler 6.2: Metrics
**Endpoint:** `GET /metrics`  
**Description:** Prometheus metrics

**TODO: CLAUDE_CODE**
```go
// TODO: CLAUDE_CODE
func (h *SystemHandler) Metrics(c *gin.Context) {
    // Use prometheus/promhttp handler
    // This is typically set up directly in routes, not as a custom handler
    // Just document that we're using:
    // promhttp.Handler()
}
```

---

## 7. Request/Response Models

### File: `internal/api/models/requests.go`

**TODO: CLAUDE_CODE**
```go
// internal/api/models/requests.go
package models

// All request models go here
// Use validation tags from go-playground/validator

// Examples:
type SendMessageRequest struct {
    Phone            string   `json:"phone" binding:"required,e164"`
    Type             string   `json:"type" binding:"required,oneof=text image document audio video template"`
    Content          string   `json:"content,omitempty" binding:"required_if=Type text"`
    MediaURL         string   `json:"media_url,omitempty" binding:"required_unless=Type text,omitempty,url"`
    Caption          string   `json:"caption,omitempty"`
    Filename         string   `json:"filename,omitempty"`
    TemplateName     string   `json:"template_name,omitempty" binding:"required_if=Type template"`
    TemplateLanguage string   `json:"template_language,omitempty" binding:"required_if=Type template"`
    Parameters       []string `json:"parameters,omitempty"`
}

// Add all other request models here
```

---

### File: `internal/api/models/responses.go`

**TODO: CLAUDE_CODE**
```go
// internal/api/models/responses.go
package models

// Standard response wrappers

type SuccessResponse struct {
    Data interface{} `json:"data,omitempty"`
}

type ListResponse struct {
    Data       interface{}        `json:"data"`
    Pagination PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
    Total   int64 `json:"total"`
    Limit   int   `json:"limit"`
    Offset  int   `json:"offset"`
    HasMore bool  `json:"has_more"`
}

type ErrorResponse struct {
    Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
    Code    string                 `json:"code"`
    Message string                 `json:"message"`
    Details map[string]interface{} `json:"details,omitempty"`
}

// Add resource-specific response models if needed
```

---

## Handler Registration

### File: `internal/api/handlers/handlers.go`

**TODO: CLAUDE_CODE**
```go
// internal/api/handlers/handlers.go
package handlers

// Container for all handlers
type Handlers struct {
    Messages  *MessagesHandler
    Contacts  *ContactsHandler
    Templates *TemplatesHandler
    Calls     *CallsHandler
    Webhooks  *WebhooksHandler
    System    *SystemHandler
}

// NewHandlers creates all handlers
func NewHandlers(
    messageService services.MessageService,
    contactService services.ContactService,
    templateService services.TemplateService,
    db *gorm.DB,
    whatsappClient *whatsapp.WhatsAppClient,
    config *config.Config,
    logger *zap.Logger,
) *Handlers {
    return &Handlers{
        Messages:  NewMessagesHandler(messageService, logger),
        Contacts:  NewContactsHandler(contactService, logger),
        Templates: NewTemplatesHandler(templateService, logger),
        Calls:     NewCallsHandler(logger), // Stub for now
        Webhooks:  NewWebhooksHandler(messageService, config.WhatsApp.WebhookVerifyToken, logger),
        System:    NewSystemHandler(db, whatsappClient, config.Server.Version, logger),
    }
}
```

---

## Testing Handlers

### Test Structure

**TODO: CLAUDE_CODE**
```go
// internal/api/handlers/messages_test.go
package handlers_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// For each handler method, test:
// 1. Success case
// 2. Validation errors
// 3. Service errors
// 4. Not found errors
// 5. Authentication failures (if applicable)

// Example:
func TestMessagesHandler_SendMessage_Success(t *testing.T) {
    // TODO: CLAUDE_CODE
    // - Create mock service
    // - Create handler with mock
    // - Create test request
    // - Call handler
    // - Assert response status and body
}

func TestMessagesHandler_SendMessage_InvalidPhone(t *testing.T) {
    // TODO: CLAUDE_CODE
    // Test validation error
}

func TestMessagesHandler_SendMessage_ServiceError(t *testing.T) {
    // TODO: CLAUDE_CODE
    // Mock service to return error
    // Assert proper error response
}
```

---

## Input Validation

**Important validation rules:**

1. **Phone Numbers:**
   - Must be E.164 format: `+[country code][number]`
   - Regex: `^\+[1-9]\d{1,14}$`
   - No spaces or special characters

2. **Message Content:**
   - Text messages: max 4096 characters
   - Captions: max 1024 characters

3. **URLs:**
   - Must be valid HTTP/HTTPS
   - Must be reachable (optional check)

4. **Enums:**
   - Use `binding:"oneof=value1 value2 value3"`
   - Validate against constants

5. **Dates:**
   - ISO 8601 format
   - Validate date ranges

---

## Error Handling in Handlers

**Pattern:**
```go
func (h *Handler) Method(c *gin.Context) {
    // Parse request
    var req RequestModel
    if err := c.ShouldBindJSON(&req); err != nil {
        // Validation error
        c.JSON(400, gin.H{
            "error": map[string]interface{}{
                "code": "invalid_request",
                "message": "Validation failed",
                "details": err.Error(),
            },
        })
        return
    }
    
    // Call service
    result, err := h.service.DoSomething(req)
    if err != nil {
        // Handle different error types
        switch e := err.(type) {
        case *errors.NotFoundError:
            c.JSON(404, gin.H{"error": e})
        case *errors.ValidationError:
            c.JSON(400, gin.H{"error": e})
        case *errors.WhatsAppError:
            c.JSON(502, gin.H{"error": e})
        default:
            c.JSON(500, gin.H{"error": "internal_error"})
        }
        return
    }
    
    // Success
    c.JSON(200, result)
}
```

---

## Documentation

**TODO: CLAUDE_CODE**

For each handler, add godoc comments:

```go
// SendMessage handles POST /api/v1/messages
// It sends a WhatsApp message (text, media, or template)
//
// Request Body:
//   - phone: recipient phone number (E.164 format)
//   - type: message type (text, image, document, etc.)
//   - content: message text (required for text messages)
//   - media_url: URL of media (required for media messages)
//
// Response: 201 Created
//   - id: message ID
//   - phone: recipient phone
//   - status: message status
//   - timestamp: when message was sent
//
// Errors:
//   - 400: invalid request (validation errors)
//   - 401: unauthorized (invalid API key)
//   - 502: WhatsApp API error
//   - 500: internal server error
func (h *MessagesHandler) SendMessage(c *gin.Context) {
    // ...
}
```

---

## Integration with Routes

**TODO: CLAUDE_CODE**
```go
// internal/api/routes/routes.go
package routes

func RegisterRoutes(router *gin.Engine, handlers *handlers.Handlers) {
    // API v1 group
    v1 := router.Group("/api/v1")
    v1.Use(middleware.AuthMiddleware())
    {
        // Messages
        messages := v1.Group("/messages")
        {
            messages.POST("", handlers.Messages.SendMessage)
            messages.GET("", handlers.Messages.ListMessages)
            messages.GET("/:id", handlers.Messages.GetMessage)
            messages.GET("/search", handlers.Messages.SearchMessages)
        }
        
        // Contacts
        contacts := v1.Group("/contacts")
        {
            contacts.GET("", handlers.Contacts.ListContacts)
            contacts.GET("/:id", handlers.Contacts.GetContact)
            contacts.PATCH("/:id", handlers.Contacts.UpdateContact)
            contacts.GET("/search", handlers.Contacts.SearchContacts)
        }
        
        // Templates
        templates := v1.Group("/templates")
        {
            templates.GET("", handlers.Templates.ListTemplates)
            templates.GET("/:id", handlers.Templates.GetTemplate)
            templates.POST("", handlers.Templates.CreateTemplate)
            templates.PATCH("/:id", handlers.Templates.UpdateTemplate)
            templates.DELETE("/:id", handlers.Templates.DeleteTemplate)
        }
        
        // Calls (future)
        calls := v1.Group("/calls")
        {
            calls.POST("", handlers.Calls.InitiateCall)
            // ... other call endpoints
        }
    }
    
    // Webhooks (no auth)
    webhooks := router.Group("/webhooks")
    webhooks.Use(middleware.WebhookAuthMiddleware(config.WebhookSecret))
    {
        webhooks.GET("/whatsapp", handlers.Webhooks.VerifyWebhook)
        webhooks.POST("/whatsapp", handlers.Webhooks.ReceiveWebhook)
    }
    
    // System (no auth)
    router.GET("/health", handlers.System.HealthCheck)
    router.GET("/metrics", promhttp.Handler())
}
```

---

## Summary

**Total Handlers to Implement:**
- ✅ Messages: 4 handlers
- ✅ Contacts: 4 handlers
- ✅ Templates: 5 handlers
- 🔲 Calls: 10 handlers (future)
- ✅ Webhooks: 2 handlers
- ✅ System: 2 handlers

**Implementation Priority:**
1. System handlers (health check) - needed for testing
2. Messages handlers - core functionality
3. Contacts handlers - needed for messages
4. Webhook handlers - needed for receiving
5. Templates handlers - extended functionality
6. Calls handlers - future phase

**Testing Coverage:**
- Each handler needs 3-5 test cases
- Focus on happy path and error cases
- Use mocks for services

---

**Status:** Ready for implementation  
**Next:** Begin with System handlers, then Messages handlers
