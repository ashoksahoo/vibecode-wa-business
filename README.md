# Vibecoded WA Client

> 🎵 A fully open-source WhatsApp Business client with messaging AND voice calling, built entirely through vibecoding

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Status](https://img.shields.io/badge/Status-Planning-yellow)](./docs/PRD.md)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)

## 🚀 What is This?

A self-hosted, production-ready WhatsApp Business API client that provides:

### Messaging (Phase 1-3)
- ✅ **REST API** for sending/receiving WhatsApp messages
- ✅ **Webhook handling** for real-time message events
- ✅ **Contact & conversation management**
- ✅ **Message history & search**
- ✅ **MCP server** for Claude AI integration
- ✅ **Docker deployment** for easy self-hosting

### Voice Calling (Phase 4) 🎙️ **NEW**
- ✅ **Inbound & outbound voice calls** via SIP/WebRTC
- ✅ **Automatic call recording** with storage
- ✅ **AI transcription** (Whisper/Deepgram)
- ✅ **Voice agent integration** (Pipecat-like)
- ✅ **Claude voice integration** via MCP
- ✅ **Call analytics** and searchable transcripts

Built entirely through **vibecoding** (AI-assisted development) to showcase how modern AI tools can help create production-quality software.

## 🎯 Status

**Current Phase:** 🟡 Planning & Documentation

We're currently finalizing the PRD and architecture before starting development.

📚 **Documentation:**
- [PRD.md](./docs/PRD.md) - Messaging features (Phases 1-3)
- [PRD_VOICE_CALLING.md](./docs/PRD_VOICE_CALLING.md) - Voice calling features (Phase 4)
- [PROJECT_INDEX.md](./docs/PROJECT_INDEX.md) - Documentation guide

## 🚫 Contribution Policy

This is a **solo-maintained project**. I am not accepting pull requests, issues, or external contributions at this time.

**However, you are free to:**
- ✅ Fork this project and build your own version
- ✅ Modify it for your needs
- ✅ Run your own instance
- ✅ Create derivative works
- ✅ Use it commercially (MIT License)

If you want to build on this project, please fork it and develop independently.

## 📋 Features

### Messaging Features (Phases 1-3)
- Send text, media, and template messages via REST API
- Receive WhatsApp messages through webhooks
- Store and query message history
- Manage contacts automatically
- Full-text message search
- PostgreSQL database for persistence

### Voice Features (Phase 4) 🎙️
- **Make & receive voice calls** on WhatsApp
- **Automatic recording** of all calls (stereo, multi-format)
- **AI transcription** with speaker diarization
- **Voice AI agents** for automated conversations
- **Searchable transcripts** across all calls
- **Call analytics** and reporting

### MCP Integration (Claude)
**Messaging Tools:**
- Send WhatsApp messages through natural language
- Query message history conversationally
- Search contacts and conversations

**Voice Tools:** 🎙️
- Make voice calls via Claude
- Retrieve call transcripts
- Search call history
- Analyze call patterns

### DevOps
- Docker containerization
- Docker Compose setup
- Environment-based configuration
- Health check endpoints
- Prometheus metrics
- S3/Minio for media storage

## 🏗️ Architecture

```
┌──────────────────┐
│   Your App       │
│  (Any Language)  │
└────────┬─────────┘
         │ REST API
┌────────▼─────────────────┐
│  Vibecoded WA Client     │
│  (Go + Gin + GORM)       │
│  ┌────────────────────┐  │
│  │  REST API          │  │
│  │  Webhook Handler   │  │
│  │  MCP Server        │  │
│  │  SIP/WebRTC 🎙️   │  │
│  │  Voice Agents 🎙️  │  │
│  └────────────────────┘  │
└────────┬────────┬────────┘
         │        │
    ┌────▼────┐  │
    │PostgreSQL│  │
    │+ Minio  │  │
    └─────────┘  │
          ┌──────▼────────────┐
          │WhatsApp Business  │
          │  Cloud API        │
          │  SIP/WebRTC 🎙️  │
          └───────────────────┘
```

## 🛠️ Tech Stack

### Core
- **Language:** Go 1.21+
- **Web Framework:** Gin
- **Database:** PostgreSQL + GORM
- **Deployment:** Docker

### Voice (Phase 4) 🎙️
- **SIP:** sipgo
- **WebRTC:** pion/webrtc
- **Codecs:** Opus
- **Transcription:** Whisper / Deepgram
- **Storage:** S3 / Minio

### APIs
- **Messaging:** WhatsApp Cloud API
- **Voice:** WhatsApp Business SIP/WebRTC

## 📖 Documentation

- **[PRD.md](./docs/PRD.md)** - Complete product requirements (messaging)
- **[PRD_VOICE_CALLING.md](./docs/PRD_VOICE_CALLING.md)** - Voice calling extension
- **[ROADMAP.md](./docs/ROADMAP.md)** - Week-by-week timeline
- **[PROJECT_INDEX.md](./docs/PROJECT_INDEX.md)** - Documentation guide
- **[CONTRIBUTING.md](./CONTRIBUTING.md)** - Fork & contribution policy
- Architecture Guide - *Coming soon*
- API Reference - *Coming soon*
- Deployment Guide - *Coming soon*

## 🚦 Roadmap

### Phase 1: Foundation (Weeks 1-2)
- [x] PRD and planning
- [ ] Core infrastructure setup
- [ ] Send/receive messages
- [ ] Webhooks

### Phase 2: Core Features (Weeks 3-4)
- [ ] Contact management
- [ ] Media messages
- [ ] Templates
- [ ] Message search

### Phase 3: MCP Server (Weeks 5-6)
- [ ] JSON-RPC implementation
- [ ] Claude integration tools
- [ ] Testing and documentation

### Phase 4: Voice Calling (Weeks 9-12) 🎙️
- [ ] SIP/WebRTC infrastructure
- [ ] Call recording
- [ ] Transcription pipeline
- [ ] Voice agent framework

### Phase 5: Polish (Weeks 7-8)
- [ ] Observability (metrics, logging)
- [ ] Rate limiting
- [ ] Advanced features
- [ ] Production readiness

See [PRD_VOICE_CALLING.md](./docs/PRD_VOICE_CALLING.md) for detailed timelines.

## 🎯 Use Cases

### For Developers
- Build WhatsApp integrations for your business
- Self-host for data privacy
- Customize for specific needs
- Integrate with existing systems

### For AI Developers
- Connect Claude to WhatsApp (messaging + voice)
- Build conversational AI on WhatsApp
- Create voice AI assistants
- Automate customer support

### For Businesses
- Send order confirmations
- Provide customer support (text + voice)
- Automated appointment reminders
- Voice-enabled customer service
- Call recording and quality monitoring

## 🎙️ Voice Calling Highlights

The voice calling features make this project unique:

- **Bidirectional calls**: Both inbound and outbound
- **Automatic recording**: Every call saved in multiple formats
- **AI transcription**: Automatic speech-to-text with speaker identification
- **Searchable history**: Find calls by searching transcript content
- **Voice AI ready**: Built-in framework for voice agents (Pipecat-like)
- **Claude integration**: Make calls and review transcripts via natural language
- **Production quality**: Opus codec, STUN/TURN support, proper SIP handling

## 📝 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

You're free to use, modify, and distribute this code. The only requirement is to include the original copyright notice.

## 🙏 Acknowledgments

- Built with [Claude](https://claude.ai) - vibecoding at its finest
- WhatsApp Business API by Meta
- Open source Go community
- Pipecat (inspiration for voice agent framework)
- pion/webrtc and sipgo libraries

## 📬 Contact

**Project Maintainer:** Ashok

For questions about the project, please check the documentation first. Remember that this is a solo project without community support channels.

---

**⭐ If you find this useful, please star the repo and share it with others!**

Built with 🎵 vibecoding | MIT License | Solo-maintained

**New:** 🎙️ Voice calling support coming in Phase 4!
