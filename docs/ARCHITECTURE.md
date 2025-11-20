# Architecture
# Vibecoded WA Client

**Last Updated:** November 18, 2025  
**Status:** 🟡 Planning Phase

---

## Overview

This document outlines the architectural decisions, system design, and technical approach for Vibecoded WA Client. All major technical decisions are documented here with rationale.

---

## Architecture Decisions

### Decision 1: WhatsApp API Approach ✅

**Decision:** WhatsApp Cloud API (hosted by Meta)

**Rationale:**
- ✅ Easiest to get started
- ✅ No infrastructure overhead
- ✅ Better documentation
- ✅ Free tier available
- ✅ Includes calling API support

**Trade-offs:**
- ⚠️ Less control vs on-premises
- ✅ But: easier for open source users

---

### Decision 2: Primary Database ✅

**Decision:** PostgreSQL

**Rationale:**
- ✅ ACID compliance
- ✅ Excellent querying
- ✅ Full-text search support
- ✅ JSON column support
- ✅ Great Go libraries (GORM)

---

### Decision 3: API Authentication ✅

**Decision:** API Keys

**Rationale:**
- ✅ Simple, stateless
- ✅ Easy for developers
- ✅ Per-key rate limiting

**Implementation:**
- Header: `Authorization: Bearer <api_key>`
- Keys stored hashed (bcrypt)
- For UI: Session-based with cookies

---

### Decision 4: Go Web Framework ✅

**Decision:** Gin

**Rationale:**
- ✅ Most popular Go framework
- ✅ Excellent performance
- ✅ Rich middleware ecosystem
- ✅ Built-in validation

---

### Decision 5: ORM ✅

**Decision:** GORM

**Rationale:**
- ✅ Migrations built-in
- ✅ Relationships handling
- ✅ Faster development
- ✅ Vibecoding friendly

---

### Decision 6: Deployment ✅

**Decision:** Docker Container

**Rationale:**
- ✅ Easy to deploy anywhere
- ✅ Reproducible builds
- ✅ Docker Compose for local dev
- ✅ Kubernetes-ready

---

### Decision 7: Configuration ✅

**Decision:** Environment Variables

**Rationale:**
- ✅ 12-factor app methodology
- ✅ Docker/Kubernetes friendly
- ✅ Secure secrets handling

---

### Decision 8: Frontend Framework

**Decision:** React 18+ or Vue 3+

**Rationale:**
- ✅ Rich ecosystem
- ✅ TypeScript support
- ✅ Component libraries

TODO: CLAUDE_CODE - Final selection

---

### Decision 9: WebRTC Signaling ✅

**Decision:** WebSocket server in Go

**Rationale:**
- ✅ Full control
- ✅ Low latency
- ✅ Integrated with main app
- ✅ No external dependencies

---

### Decision 10: Call Recording Storage ✅

**Decision:** Local filesystem (with S3 option)

**Rationale:**
- ✅ Simple
- ✅ Fast access
- ✅ Low cost
- ✅ Can migrate to S3 later

---

### Decision 11: Message Queue ✅

**Decision:** In-memory (Go channels) with optional Redis

**Rationale:**
- ✅ No external dependency initially
- ✅ Simple, fast
- ✅ Can add Redis later

---

### Decision 12: Transcription Service ✅

**Decision:** Deepgram

**Rationale:**
- ✅ Excellent accuracy
- ✅ Low latency streaming
- ✅ Speaker diarization
- ✅ Multiple languages

---

### Decision 13: Text-to-Speech ✅

**Decision:** 11labs

**Rationale:**
- ✅ Most natural voices
- ✅ Custom voice cloning
- ✅ Emotion control
- ✅ Good for IVR

---

## System Architecture

### High-Level Overview

```
Users → REST API / WebSocket / MCP Server
           ↓
    Business Logic Layer
           ↓
    Data Access Layer (GORM)
           ↓
    PostgreSQL Database

External Services:
- WhatsApp Cloud API
- Deepgram (Transcription)
- 11labs (TTS)
```

---

## Component Architecture

### API Layer
- REST API (Gin)
- WebSocket Server (WebRTC signaling)
- MCP Server (JSON-RPC)

### Service Layer
- Message Service
- Contact Service
- Template Service
- Call Service
- Transcription Service (Deepgram)
- TTS Service (11labs)
- WhatsApp Client Service

### Data Access Layer
- Repository Pattern
- MessageRepository
- ContactRepository
- CallRepository, etc.

---

## Data Flow

### Send Message
```
User → API → Message Service → WhatsApp API
        ↓
   Database (save)
        ↓
   Return message ID
```

### Receive Message
```
WhatsApp → Webhook → Verify → Parse → Save → Update Contact
```

### Call Flow
```
User → Call Service → WhatsApp API → WebRTC
        ↓                              ↓
   Deepgram (transcribe)          11labs (TTS)
        ↓                              ↓
   Save transcription            Voice response
```

---

## Deployment Architecture

### Development
- Docker Compose
- App container + PostgreSQL container
- Local volumes
- .env file

### Production
- Docker container on cloud (AWS/GCP/Azure/DO)
- Separate PostgreSQL database
- Persistent volumes for recordings
- HTTPS/WSS via load balancer
- Environment variables from secrets

---

**Back:** [Technical Requirements](TECHNICAL_REQUIREMENTS.md) | **Next:** [API Design](API_DESIGN.md)
