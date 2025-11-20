# Vision & Product Overview
# Vibecoded WA Client

**Version:** 0.1.0  
**Last Updated:** November 18, 2025

---

## Executive Summary

**Project Name:** Vibecoded WA Client  
**Type:** Open Source WhatsApp Business Client + MCP Server  
**Primary Language:** Go  
**Development Approach:** Vibecoding (AI-assisted development)  
**Contribution Policy:** Solo-maintained, no external contributions (fork-friendly)

### What We're Building
A self-hosted, open-source WhatsApp Business client that provides:
- RESTful API for sending/receiving WhatsApp messages
- Webhook handling for inbound messages
- Contact and conversation management
- Message history and search
- User management and API key system
- MCP server for Claude integration
- Docker deployment support

### Why This Exists
To provide developers with a fully customizable, self-hosted WhatsApp Business solution that can be integrated into any system, while showcasing the effectiveness of vibecoding methodology.

---

## Vision Statement

Build a production-ready WhatsApp Business client that developers can deploy, customize, and integrate freely. Demonstrate that complex backend systems can be built entirely through vibecoding while maintaining code quality and documentation standards.

---

## Primary Goals

- ✅ Full WhatsApp Business API integration
- ✅ RESTful API for easy integration
- ✅ Webhook support for real-time message handling
- ✅ Contact and conversation management
- ✅ User management with role-based permissions
- ✅ MCP server for Claude AI integration
- ✅ Self-hosted deployment (Docker)
- ✅ Comprehensive documentation
- ✅ MIT licensed and fork-friendly

---

## Non-Goals

- ❌ Building a UI/frontend (API-only)
- ❌ Supporting personal WhatsApp accounts
- ❌ Multi-tenant SaaS architecture
- ❌ Accepting pull requests (solo project)
- ❌ Commercial support or SLA guarantees
- ❌ Unofficial WhatsApp protocol reverse engineering

---

## Target Users

### Primary Persona: Backend Developer
**Profile:**
- Mid to senior-level developer
- Building business automation systems
- Needs WhatsApp integration for customer communication
- Wants control over infrastructure

**Pain Points:**
- WhatsApp's official API is complex to integrate
- Need more flexibility than official SDKs provide
- Want to self-host for data privacy
- Existing solutions are expensive or limited

**Use Cases:**
- E-commerce order notifications
- Customer support chatbots
- Appointment reminders
- Marketing campaigns
- Internal team communication

### Secondary Persona: AI Developer
**Profile:**
- Working with Claude or other LLMs
- Building conversational AI systems
- Needs WhatsApp as a communication channel

**Pain Points:**
- No easy way to connect Claude to WhatsApp
- Want natural language control of messaging
- Need bidirectional communication

**Use Cases:**
- AI assistants on WhatsApp
- Automated customer service
- Personal productivity bots
- Research projects

### Tertiary Persona: Solo Developer/Indie Hacker
**Profile:**
- Building side projects or MVPs
- Budget-conscious
- Wants simple deployment

**Pain Points:**
- Commercial solutions too expensive
- Need something they can modify
- Want to learn and experiment

**Use Cases:**
- Personal projects
- Learning WhatsApp API
- Building custom integrations

---

## Success Criteria

### Technical Metrics
- 📊 API response time < 200ms (p95)
- 📊 Message delivery success > 99%
- 📊 Webhook processing < 1s
- 📊 Database query time < 50ms (p95)
- 📊 Zero data loss
- 📊 System uptime > 99.5%

### User Metrics
- 📊 Time to send first message < 10 minutes
- 📊 Documentation comprehensiveness: Can deploy without asking questions
- 📊 GitHub stars > 100 (6 months)
- 📊 Active forks > 20 (6 months)

### Project Health
- ✅ README is comprehensive
- ✅ All features documented
- ✅ API reference complete
- ✅ Docker deployment tested
- ✅ Example code provided
