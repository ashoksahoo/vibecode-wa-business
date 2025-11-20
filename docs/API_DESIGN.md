# API Design
# Vibecoded WA Client

**Last Updated:** November 18, 2025  
**Status:** 🟡 Planning Phase

---

## Overview

This document specifies the complete REST API design for Vibecoded WA Client, including all endpoints, request/response formats, authentication, and error handling.

---

## API Endpoints Index

### Messages API
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/messages` | Send text, media, or template message |
| GET | `/api/v1/messages` | List messages with filters |
| GET | `/api/v1/messages/:id` | Get specific message |
| GET | `/api/v1/messages/search` | Search messages by content |

### Contacts API
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/contacts` | List all contacts |
| GET | `/api/v1/contacts/:id` | Get contact details |
| PATCH | `/api/v1/contacts/:id` | Update contact |
| GET | `/api/v1/contacts/search` | Search contacts |

### Templates API
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/templates` | List all templates |
| GET | `/api/v1/templates/:id` | Get template details |
| POST | `/api/v1/templates` | Create new template |
| PATCH | `/api/v1/templates/:id` | Update template |
| DELETE | `/api/v1/templates/:id` | Delete template |

### Calls API
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/calls` | Initiate outbound call |
| GET | `/api/v1/calls` | List call history |
| GET | `/api/v1/calls/:id` | Get call details |
| DELETE | `/api/v1/calls/:id` | End active call |
| POST | `/api/v1/calls/:id/record` | Start call recording |
| DELETE | `/api/v1/calls/:id/record` | Stop call recording |
| GET | `/api/v1/calls/:id/recording` | Get call recording |
| GET | `/api/v1/calls/:id/transcription` | Get call transcription |
| POST | `/api/v1/calls/:id/speak` | Send TTS during call |
| PATCH | `/api/v1/calls/:id` | Call control (mute, hold, transfer) |

### Webhooks
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/webhooks/whatsapp` | Receive WhatsApp events |
| GET | `/webhooks/whatsapp` | Webhook verification |

### System API
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/metrics` | Prometheus metrics |

**Total Endpoints:** 27

---

## API Endpoints Index

### Messages API
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/messages` | Send text, media, or template message |
| GET | `/api/v1/messages` | List messages with filters |
| GET | `/api/v1/messages/:id` | Get specific message |
| GET | `/api/v1/messages/search` | Search messages by content |

### Contacts API
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/contacts` | List all contacts |
| GET | `/api/v1/contacts/:id` | Get contact details |
| PATCH | `/api/v1/contacts/:id` | Update contact |
| GET | `/api/v1/contacts/search` | Search contacts |

### Templates API
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/templates` | List all templates |
| GET | `/api/v1/templates/:id` | Get template details |
| POST | `/api/v1/templates` | Create new template |
| PATCH | `/api/v1/templates/:id` | Update template |
| DELETE | `/api/v1/templates/:id` | Delete template |

### Calls API
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/calls` | Initiate outbound call |
| GET | `/api/v1/calls` | List call history |
| GET | `/api/v1/calls/:id` | Get call details |
| DELETE | `/api/v1/calls/:id` | End active call |
| POST | `/api/v1/calls/:id/record` | Start call recording |
| DELETE | `/api/v1/calls/:id/record` | Stop call recording |
| GET | `/api/v1/calls/:id/recording` | Get call recording |
| GET | `/api/v1/calls/:id/transcription` | Get call transcription |
| POST | `/api/v1/calls/:id/speak` | Send TTS during call |
| PATCH | `/api/v1/calls/:id` | Call control (mute, hold, transfer) |

### Webhooks
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/webhooks/whatsapp` | Receive WhatsApp events |
| GET | `/webhooks/whatsapp` | Webhook verification |

### System API
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/metrics` | Prometheus metrics |

**Total Endpoints:** 27

---

## Base Configuration

### Base URL
```
http://localhost:8080/api/v1
```

Production:
```
https://your-domain.com/api/v1
```

### Authentication

All API endpoints (except webhooks and health) require authentication:

```http
Authorization: Bearer <api_key>
```

### Content Type
```http
Content-Type: application/json
```

### Response Format

**Success Response:**
```json
{
  "id": "resource_id",
  "field": "value",
  "created_at": "2025-11-18T10:30:00Z"
}
```

**List Response:**
```json
{
  "data": [...],
  "pagination": {
    "total": 150,
    "limit": 20,
    "offset": 40,
    "has_more": true
  }
}
```

**Error Response:**
```json
{
  "error": {
    "code": "error_code",
    "message": "Human-readable error message",
    "details": {
      "field": "additional_info"
    }
  }
}
```

---

## Messages API

### Send Message

**Endpoint:** `POST /api/v1/messages`

**Description:** Send text, media, or template messages to WhatsApp

**Common Parameters:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| phone | string | Yes | Phone number in international format (E.164) |
| type | string | Yes | Message type: "text", "image", "document", "audio", "video", "template" |

**Type-specific Parameters:**

**For Text Messages:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| content | string | Yes | Message text (max 4096 chars) |

**For Media Messages (image, document, audio, video):**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| media_url | string | Yes | URL of media file (must be publicly accessible) |
| caption | string | No | Media caption (max 1024 chars) |
| filename | string | No | Filename (for documents) |

**For Template Messages:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| template_name | string | Yes | Template name from WhatsApp |
| template_language | string | Yes | Language code (e.g., "en", "es") |
| parameters | array | No | Array of parameter values |

---

**Example Requests:**

**Text Message:**
```json
{
  "phone": "+1234567890",
  "type": "text",
  "content": "Hello from Vibecoded WA Client!"
}
```

**Image Message:**
```json
{
  "phone": "+1234567890",
  "type": "image",
  "media_url": "https://example.com/image.jpg",
  "caption": "Check this out!"
}
```

**Document Message:**
```json
{
  "phone": "+1234567890",
  "type": "document",
  "media_url": "https://example.com/invoice.pdf",
  "filename": "invoice.pdf",
  "caption": "Your invoice"
}
```

**Template Message:**
```json
{
  "phone": "+1234567890",
  "type": "template",
  "template_name": "order_confirmation",
  "template_language": "en",
  "parameters": ["12345", "John Doe", "$99.99"]
}
```

---

**Response:** `201 Created`
```json
{
  "id": "msg_abc123",
  "phone": "+1234567890",
  "type": "text",
  "content": "Hello from Vibecoded WA Client!",
  "status": "sent",
  "direction": "outbound",
  "whatsapp_message_id": "wamid.HBgLMTIzNDU2Nzg5MAo...",
  "created_at": "2025-11-18T10:30:00Z"
}
```

**Supported Media Types:**
- **Image:** JPEG, PNG (max 5MB)
- **Document:** PDF, DOC, DOCX, XLS, XLSX, etc. (max 100MB)
- **Audio:** AAC, MP3, OGG (max 16MB)
- **Video:** MP4, 3GP (max 16MB)

**Errors:**
- `400` - Invalid phone number format
- `400` - Content too long
- `400` - Invalid media URL
- `400` - Template not found
- `401` - Invalid API key
- `429` - Rate limit exceeded
- `502` - WhatsApp API error

---

### Get Message by ID

**Endpoint:** `GET /api/v1/messages/:id`

**Description:** Retrieve a specific message

**Response:** `200 OK`
```json
{
  "id": "msg_abc123",
  "phone": "+1234567890",
  "direction": "outbound",
  "type": "text",
  "content": "Hello!",
  "status": "delivered",
  "whatsapp_message_id": "wamid.HBgLMTIzNDU2Nzg5MAo...",
  "created_at": "2025-11-18T10:30:00Z",
  "updated_at": "2025-11-18T10:30:15Z"
}
```

**Errors:**
- `404` - Message not found

---

### List Messages

**Endpoint:** `GET /api/v1/messages`

**Description:** List messages with pagination and filters

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| phone | string | No | Filter by phone number |
| direction | string | No | "inbound" or "outbound" |
| type | string | No | Message type filter |
| status | string | No | Status filter |
| start_date | string | No | ISO 8601 date |
| end_date | string | No | ISO 8601 date |
| limit | integer | No | Results per page (default: 20, max: 100) |
| offset | integer | No | Pagination offset (default: 0) |

**Example Request:**
```
GET /api/v1/messages?phone=+1234567890&limit=20&offset=0
```

**Response:** `200 OK`
```json
{
  "data": [
    {
      "id": "msg_abc123",
      "phone": "+1234567890",
      "direction": "outbound",
      "type": "text",
      "content": "Hello!",
      "status": "delivered",
      "created_at": "2025-11-18T10:30:00Z"
    },
    {
      "id": "msg_abc122",
      "phone": "+1234567890",
      "direction": "inbound",
      "type": "text",
      "content": "Hi there!",
      "status": "received",
      "created_at": "2025-11-18T10:29:00Z"
    }
  ],
  "pagination": {
    "total": 150,
    "limit": 20,
    "offset": 0,
    "has_more": true
  }
}
```

---

### Search Messages

**Endpoint:** `GET /api/v1/messages/search`

**Description:** Full-text search in message content

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| q | string | Yes | Search query |
| phone | string | No | Limit search to specific phone |
| limit | integer | No | Results limit (default: 20) |

**Example Request:**
```
GET /api/v1/messages/search?q=order&phone=+1234567890
```

**Response:** `200 OK`
```json
{
  "data": [
    {
      "id": "msg_abc120",
      "phone": "+1234567890",
      "content": "Your order #12345 has been confirmed",
      "created_at": "2025-11-18T09:00:00Z"
    }
  ],
  "count": 1
}
```

---

## Contacts API

### List Contacts

**Endpoint:** `GET /api/v1/contacts`

**Description:** List all contacts with pagination

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| limit | integer | No | Results per page (default: 50, max: 100) |
| offset | integer | No | Pagination offset |
| sort | string | No | Sort field (default: "last_message_at") |
| order | string | No | "asc" or "desc" (default: "desc") |

**Response:** `200 OK`
```json
{
  "data": [
    {
      "id": "contact_xyz789",
      "phone": "+1234567890",
      "name": "John Doe",
      "profile_picture_url": "https://...",
      "last_message_at": "2025-11-18T10:30:00Z",
      "message_count": 45,
      "unread_count": 2,
      "created_at": "2025-01-15T08:00:00Z"
    }
  ],
  "pagination": {
    "total": 120,
    "limit": 50,
    "offset": 0,
    "has_more": true
  }
}
```

---

### Get Contact

**Endpoint:** `GET /api/v1/contacts/:id`

**Description:** Get contact details

**Response:** `200 OK`
```json
{
  "id": "contact_xyz789",
  "phone": "+1234567890",
  "name": "John Doe",
  "profile_picture_url": "https://...",
  "last_message_at": "2025-11-18T10:30:00Z",
  "message_count": 45,
  "unread_count": 2,
  "metadata": {
    "customer_id": "CUS123",
    "tags": ["vip", "support"]
  },
  "created_at": "2025-01-15T08:00:00Z",
  "updated_at": "2025-11-18T10:30:00Z"
}
```

---

### Update Contact

**Endpoint:** `PATCH /api/v1/contacts/:id`

**Description:** Update contact information

**Request Body:**
```json
{
  "name": "John Smith",
  "metadata": {
    "customer_id": "CUS123",
    "tags": ["vip", "enterprise"]
  }
}
```

**Response:** `200 OK`
```json
{
  "id": "contact_xyz789",
  "phone": "+1234567890",
  "name": "John Smith",
  "metadata": {
    "customer_id": "CUS123",
    "tags": ["vip", "enterprise"]
  },
  "updated_at": "2025-11-18T11:00:00Z"
}
```

---

### Search Contacts

**Endpoint:** `GET /api/v1/contacts/search`

**Description:** Search contacts by name or phone

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| q | string | Yes | Search query |
| limit | integer | No | Results limit |

**Example:**
```
GET /api/v1/contacts/search?q=john
```

**Response:** `200 OK`
```json
{
  "data": [
    {
      "id": "contact_xyz789",
      "phone": "+1234567890",
      "name": "John Doe"
    }
  ],
  "count": 1
}
```

---

## Templates API

### List Templates

**Endpoint:** `GET /api/v1/templates`

**Description:** List all message templates

**Response:** `200 OK`
```json
{
  "data": [
    {
      "id": "tmpl_abc123",
      "name": "order_confirmation",
      "language": "en",
      "category": "utility",
      "status": "approved",
      "content": "Your order {{1}} has been confirmed for {{2}}",
      "parameters": ["order_id", "customer_name"],
      "created_at": "2025-01-10T00:00:00Z"
    }
  ]
}
```

---

### Get Template

**Endpoint:** `GET /api/v1/templates/:id`

**Description:** Get template details

**Response:** `200 OK`

---

### Create Template

**Endpoint:** `POST /api/v1/templates`

**Description:** Create new template

**Request Body:**
```json
{
  "name": "appointment_reminder",
  "language": "en",
  "category": "utility",
  "content": "Hi {{1}}, your appointment is on {{2}} at {{3}}",
  "parameters": ["name", "date", "time"]
}
```

**Response:** `201 Created`

---

### Update Template

**Endpoint:** `PATCH /api/v1/templates/:id`

**Description:** Update template

---

### Delete Template

**Endpoint:** `DELETE /api/v1/templates/:id`

**Description:** Delete template

**Response:** `204 No Content`

---

## Calls API

### Initiate Call

**Endpoint:** `POST /api/v1/calls`

**Description:** Start an outbound WhatsApp call

**Request Body:**
```json
{
  "phone": "+1234567890",
  "auto_record": true
}
```

**Parameters:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| phone | string | Yes | Phone number to call |
| auto_record | boolean | No | Auto-start recording (default: false) |

**Response:** `201 Created`
```json
{
  "id": "call_xyz123",
  "phone": "+1234567890",
  "direction": "outbound",
  "status": "initiated",
  "webrtc_offer": "...",
  "created_at": "2025-11-18T10:30:00Z"
}
```

**Call Status Values:**
- `initiated` - Call started
- `ringing` - Phone is ringing
- `in-progress` - Call is active
- `completed` - Call ended normally
- `failed` - Call failed
- `rejected` - Call was rejected
- `missed` - Call not answered

---

### List Calls

**Endpoint:** `GET /api/v1/calls`

**Description:** List call history

**Query Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| phone | string | No | Filter by phone |
| direction | string | No | "inbound" or "outbound" |
| status | string | No | Filter by status |
| start_date | string | No | ISO 8601 date |
| end_date | string | No | ISO 8601 date |
| limit | integer | No | Results per page (default: 20) |
| offset | integer | No | Pagination offset |

**Response:** `200 OK`
```json
{
  "data": [
    {
      "id": "call_xyz123",
      "phone": "+1234567890",
      "direction": "outbound",
      "status": "completed",
      "duration": 125,
      "recording_url": "https://.../recording.mp3",
      "transcription_available": true,
      "created_at": "2025-11-18T10:30:00Z",
      "ended_at": "2025-11-18T10:32:05Z"
    }
  ],
  "pagination": {
    "total": 50,
    "limit": 20,
    "offset": 0,
    "has_more": true
  }
}
```

---

### Get Call Details

**Endpoint:** `GET /api/v1/calls/:id`

**Description:** Get detailed call information

**Response:** `200 OK`
```json
{
  "id": "call_xyz123",
  "phone": "+1234567890",
  "direction": "outbound",
  "status": "completed",
  "duration": 125,
  "recording_url": "https://.../recording.mp3",
  "transcription_id": "trans_abc456",
  "created_at": "2025-11-18T10:30:00Z",
  "started_at": "2025-11-18T10:30:10Z",
  "ended_at": "2025-11-18T10:32:15Z"
}
```

---

### End Call

**Endpoint:** `DELETE /api/v1/calls/:id`

**Description:** End an active call

**Response:** `200 OK`
```json
{
  "id": "call_xyz123",
  "status": "completed",
  "duration": 125,
  "ended_at": "2025-11-18T10:32:15Z"
}
```

---

### Start Call Recording

**Endpoint:** `POST /api/v1/calls/:id/record`

**Description:** Start recording during call

**Response:** `200 OK`
```json
{
  "call_id": "call_xyz123",
  "recording_started": true,
  "started_at": "2025-11-18T10:31:00Z"
}
```

---

### Stop Call Recording

**Endpoint:** `DELETE /api/v1/calls/:id/record`

**Description:** Stop recording

**Response:** `200 OK`
```json
{
  "call_id": "call_xyz123",
  "recording_stopped": true,
  "recording_url": "https://.../recording.mp3",
  "duration": 60
}
```

---

### Get Call Recording

**Endpoint:** `GET /api/v1/calls/:id/recording`

**Description:** Get recording URL and metadata

**Response:** `200 OK`
```json
{
  "call_id": "call_xyz123",
  "recording_url": "https://.../recording.mp3",
  "format": "mp3",
  "size": 1245678,
  "duration": 125,
  "created_at": "2025-11-18T10:32:15Z"
}
```

---

### Get Call Transcription

**Endpoint:** `GET /api/v1/calls/:id/transcription`

**Description:** Get call transcription

**Response:** `200 OK`
```json
{
  "call_id": "call_xyz123",
  "transcription": [
    {
      "speaker": "agent",
      "text": "Hello, how can I help you?",
      "timestamp": 0.5,
      "confidence": 0.98
    },
    {
      "speaker": "customer",
      "text": "I'd like to check my order status",
      "timestamp": 3.2,
      "confidence": 0.95
    }
  ],
  "language": "en",
  "duration": 125,
  "created_at": "2025-11-18T10:32:20Z"
}
```

---

### Send TTS During Call

**Endpoint:** `POST /api/v1/calls/:id/speak`

**Description:** Send text-to-speech audio during active call

**Request Body:**
```json
{
  "text": "Your order has been confirmed",
  "voice": "default",
  "language": "en"
}
```

**Response:** `200 OK`
```json
{
  "call_id": "call_xyz123",
  "text": "Your order has been confirmed",
  "audio_url": "https://.../tts_audio.mp3",
  "played": true
}
```

---

### Call Control Actions

**Endpoint:** `PATCH /api/v1/calls/:id`

**Description:** Control call (mute, hold, transfer)

**Request Body (Mute):**
```json
{
  "action": "mute"
}
```

**Request Body (Hold):**
```json
{
  "action": "hold"
}
```

**Request Body (Transfer):**
```json
{
  "action": "transfer",
  "target_phone": "+9876543210"
}
```

**Response:** `200 OK`
```json
{
  "call_id": "call_xyz123",
  "action": "mute",
  "success": true
}
```

---

## Webhooks

### WhatsApp Webhook (Receive)

**Endpoint:** `POST /webhooks/whatsapp`

**Description:** Receive WhatsApp events (messages, status updates, call events)

**Authentication:** Signature verification (X-Hub-Signature-256)

**Webhook Payload (Incoming Message):**
```json
{
  "object": "whatsapp_business_account",
  "entry": [{
    "changes": [{
      "value": {
        "messaging_product": "whatsapp",
        "messages": [{
          "from": "+1234567890",
          "id": "wamid.xxx",
          "timestamp": "1700000000",
          "type": "text",
          "text": {
            "body": "Hello"
          }
        }]
      }
    }]
  }]
}
```

**Webhook Payload (Message Status):**
```json
{
  "object": "whatsapp_business_account",
  "entry": [{
    "changes": [{
      "value": {
        "messaging_product": "whatsapp",
        "statuses": [{
          "id": "wamid.xxx",
          "status": "delivered",
          "timestamp": "1700000010"
        }]
      }
    }]
  }]
}
```

**Webhook Payload (Call Event):**
```json
{
  "object": "whatsapp_business_account",
  "entry": [{
    "changes": [{
      "value": {
        "messaging_product": "whatsapp",
        "call_events": [{
          "call_id": "xxx",
          "from": "+1234567890",
          "event": "ringing",
          "timestamp": "1700000000"
        }]
      }
    }]
  }]
}
```

**Response:** `200 OK`
```json
{
  "status": "received"
}
```

---

### WhatsApp Webhook Verification (GET)

**Endpoint:** `GET /webhooks/whatsapp`

**Description:** WhatsApp webhook verification

**Query Parameters:**
- `hub.mode` - Should be "subscribe"
- `hub.challenge` - Random string from WhatsApp
- `hub.verify_token` - Your verification token

**Response:** Return `hub.challenge` value

---

## System API

### Health Check

**Endpoint:** `GET /health`

**Description:** Check system health

**Authentication:** None required

**Response:** `200 OK`
```json
{
  "status": "healthy",
  "version": "0.1.0",
  "uptime": "24h30m15s",
  "checks": {
    "database": "connected",
    "whatsapp_api": "reachable",
    "deepgram_api": "reachable",
    "elevenlabs_api": "reachable"
  },
  "timestamp": "2025-11-18T10:30:00Z"
}
```

**Unhealthy Response:** `503 Service Unavailable`
```json
{
  "status": "unhealthy",
  "checks": {
    "database": "error: connection refused",
    "whatsapp_api": "reachable"
  }
}
```

---

### Metrics (Prometheus)

**Endpoint:** `GET /metrics`

**Description:** Prometheus metrics endpoint

**Authentication:** None required (or optional basic auth)

**Response:** Prometheus text format
```
# HELP http_requests_total Total number of HTTP requests
# TYPE http_requests_total counter
http_requests_total{method="POST",endpoint="/api/v1/messages",status="200"} 1543

# HELP message_send_duration_seconds Message send duration
# TYPE message_send_duration_seconds histogram
message_send_duration_seconds_bucket{le="0.1"} 850
message_send_duration_seconds_bucket{le="0.5"} 1200
```

---

## Error Codes

### HTTP Status Codes

| Code | Description |
|------|-------------|
| 200 | OK - Request successful |
| 201 | Created - Resource created |
| 204 | No Content - Successful deletion |
| 400 | Bad Request - Invalid input |
| 401 | Unauthorized - Invalid API key |
| 403 | Forbidden - Insufficient permissions |
| 404 | Not Found - Resource doesn't exist |
| 429 | Too Many Requests - Rate limit exceeded |
| 500 | Internal Server Error - Server error |
| 502 | Bad Gateway - WhatsApp API error |
| 503 | Service Unavailable - System down |

### Application Error Codes

| Code | Description |
|------|-------------|
| `invalid_phone_number` | Phone number format invalid |
| `invalid_media_url` | Media URL not accessible |
| `invalid_template` | Template not found or invalid |
| `message_too_long` | Message exceeds length limit |
| `rate_limit_exceeded` | Too many requests |
| `whatsapp_api_error` | WhatsApp API returned error |
| `call_not_found` | Call ID doesn't exist |
| `call_already_ended` | Cannot modify ended call |
| `recording_not_available` | Recording not ready |
| `transcription_failed` | Transcription error |
| `tts_failed` | Text-to-speech error |
| `internal_error` | Generic server error |

---

## Rate Limiting

### Rate Limit Headers

Every API response includes rate limit headers:

```http
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 995
X-RateLimit-Reset: 1700000000
```

### Default Limits

| Endpoint Type | Limit |
|---------------|-------|
| Default | 1000 requests/hour |
| Message sending | 100 messages/minute |
| Call initiation | 10 calls/minute |

### Rate Limit Exceeded Response

**Status:** `429 Too Many Requests`

```json
{
  "error": {
    "code": "rate_limit_exceeded",
    "message": "Rate limit exceeded. Try again in 60 seconds",
    "details": {
      "limit": 1000,
      "reset_at": "2025-11-18T11:00:00Z"
    }
  }
}
```

---

## Pagination

### Query Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| limit | integer | 20 | Items per page (max: 100) |
| offset | integer | 0 | Number of items to skip |

### Response Format

```json
{
  "data": [...],
  "pagination": {
    "total": 150,
    "limit": 20,
    "offset": 40,
    "has_more": true
  }
}
```

---

## Webhooks Security

### Signature Verification

WhatsApp sends webhooks with signature in header:
```http
X-Hub-Signature-256: sha256=<signature>
```

**Verification Algorithm:**
1. Get raw request body
2. Compute HMAC-SHA256 with webhook secret
3. Compare with signature from header
4. Reject if mismatch

**Example (pseudo-code):**
```
signature = HMAC-SHA256(webhook_secret, request_body)
expected = "sha256=" + hex(signature)
if expected != header_signature:
    reject()
```

---

## API Versioning

Current version: **v1**

Base URL includes version:
```
/api/v1/...
```

Future versions:
```
/api/v2/...
```

Breaking changes will increment version number.

---

## TODO: CLAUDE_CODE

The following need implementation:
- [ ] All endpoint handlers
- [ ] Request validation middleware
- [ ] Authentication middleware
- [ ] Rate limiting middleware
- [ ] Webhook signature verification
- [ ] Error handling middleware
- [ ] Response formatting helpers
- [ ] OpenAPI/Swagger documentation generation

---

**Back:** [Architecture](ARCHITECTURE.md) | **Next:** [Data Models](DATA_MODELS.md)
