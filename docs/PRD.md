# Product Requirements Document
# Vibecoded WA Client

**Version:** 0.1.0  
**Author:** Ashok  
**Date:** November 18, 2025  

---

## Executive Summary

**Vibecoded WA Client** is a self-hosted, open-source WhatsApp Business client providing comprehensive messaging and calling capabilities through REST APIs, WebRTC, and Claude AI integration (MCP).

**What We're Building:**
- RESTful API for WhatsApp messaging and calling
- WebRTC client for browser-based calls
- AI-powered call automation (Deepgram + 11labs)
- Web UI dashboard
- MCP server for Claude integration
- Docker deployment

**Development Approach:** Vibecoding (AI-assisted development)  
**License:** MIT  
**Contribution Policy:** Solo-maintained, fork-friendly

---

## Documentation Structure

This PRD is organized into focused documents:

### 1. [Vision & Goals](VISION.md)
Vision statement, project goals, target users, and non-goals.

### 2. [Features](FEATURES.md)
Complete feature list including:
- Messaging (text, media, templates, interactive)
- Contacts management
- WhatsApp Business calling
- WebRTC client
- Call recording & transcription
- AI automation (Deepgram/11labs)
- Web UI dashboard
- MCP server integration
- Analytics & reporting

### 3. [Technical Requirements](TECHNICAL_REQUIREMENTS.md)
Performance targets, scalability, security, testing, and compliance requirements.

### 4. [Architecture](ARCHITECTURE.md)
All architectural decisions with rationale:
- Technology choices (Go, PostgreSQL, React/Vue)
- System design
- Component architecture
- Data flow
- Deployment architecture

### 5. [API Design](API_DESIGN.md)
Complete REST API specification:
- Messages API
- Contacts API
- Templates API
- Calls API
- Webhooks
- System endpoints
- Error responses

### 6. [Data Models](DATA_MODELS.md)
Database schema and GORM models:
- Messages, Contacts, Templates
- Calls, Transcriptions
- API Keys
- Relationships and indexes

### 7. [Tasks](TASKS.md)
Breakdown of implementation tasks with TODO markers for Claude Code

---

## Quick Start

**For Developers:**
1. Read [Vision & Goals](VISION.md) to understand the project
2. Review [Features](FEATURES.md) for what's being built
3. Check [Technical Requirements](TECHNICAL_REQUIREMENTS.md) for specs
4. Study [Architecture](ARCHITECTURE.md) for design decisions
5. Reference [API Design](API_DESIGN.md) for endpoints
6. See [Tasks](TASKS.md) for implementation breakdown

**For Implementation:**
- All TODO markers indicate code to be written
- Follow architecture decisions documented
- Refer to data models for database design
- Use tech stack for library versions

---

## Key Decisions

### What This Is
✅ Self-hosted WhatsApp Business client  
✅ REST API for integration  
✅ WebRTC calling with AI automation  
✅ Web UI for management  
✅ MCP server for Claude  
✅ Docker deployment  

### What This Is Not
❌ Multi-tenant SaaS  
❌ Personal WhatsApp client  
❌ Accepting external contributions (fork-friendly)  
❌ Commercial support  

---

## Core Technologies

- **Backend:** Go 1.21+, Gin, GORM, PostgreSQL
- **Frontend:** React/Vue, TypeScript, WebRTC
- **External APIs:** WhatsApp Cloud API, Deepgram, 11labs
- **Deployment:** Docker, Docker Compose
- **Integration:** MCP (Model Context Protocol)

---

## Success Criteria

✅ Full WhatsApp messaging support  
✅ WebRTC calling with recording  
✅ AI call automation working  
✅ Web UI functional  
✅ MCP server integrated with Claude  
✅ Docker deployment < 15 minutes  
✅ API response time < 200ms (p95)  
✅ Comprehensive documentation  

---

## Getting Help

- **Documentation:** See links above
- **Tasks:** [TASKS.md](TASKS.md)
- **Issues:** GitHub Issues (read-only tracking)

---

## Document History

| Date | Version | Changes |
|------|---------|---------|
| 2025-11-18 | 0.1.0 | Initial PRD structure |
| 2025-11-18 | 0.1.1 | Split into focused documents |
| 2025-11-21 | 0.1.2 | Removed project management elements |

---

**Maintained by:** Ashok  
**Vibecoded:** Yes 🎵
