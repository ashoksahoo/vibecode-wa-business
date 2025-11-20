# Project Structure

This document describes the directory structure of the Vibecoded WA Client project.

## Overview

```
vibecoded-wa-client/
├── cmd/                    # Application entry points
│   └── server/            # Main server application
│       └── main.go        # Application entry point
│
├── internal/              # Private application code
│   ├── api/              # REST API implementation
│   │   └── server.go     # API server setup
│   ├── database/         # Database layer
│   │   └── database.go   # DB connection & migrations
│   ├── models/           # Data models
│   │   ├── message.go    # Message & contact models
│   │   └── call.go       # Voice calling models (Phase 4)
│   ├── whatsapp/         # WhatsApp API integration
│   │   └── client.go     # WhatsApp client
│   ├── mcp/              # Model Context Protocol server
│   │   └── server.go     # MCP implementation for Claude
│   └── voice/            # Voice calling (Phase 4)
│       └── sip.go        # SIP/WebRTC implementation
│
├── pkg/                   # Public/reusable packages
│   ├── config/           # Configuration management
│   │   ├── config.go     # Config loader
│   │   └── config_test.go # Config tests
│   └── logger/           # Logging utilities
│       └── logger.go     # Structured logger
│
├── migrations/            # Database migrations
├── tests/                # Test files
│   ├── integration/      # Integration tests
│   └── unit/            # Unit tests
│
├── storage/              # Local storage
│   ├── media/           # WhatsApp media files
│   └── recordings/      # Call recordings (Phase 4)
│
├── scripts/              # Build & deployment scripts
├── deployments/          # Deployment configurations
│   └── prometheus.yml   # Prometheus config
│
├── docs/                 # Documentation
│   ├── PRD.md           # Product requirements (messaging)
│   ├── PRD_VOICE_CALLING.md  # Voice calling extension
│   ├── ROADMAP.md       # Development timeline
│   └── PROJECT_INDEX.md # Documentation index
│
├── .env.example          # Environment variables template
├── .gitignore           # Git ignore rules
├── .dockerignore        # Docker ignore rules
├── Dockerfile           # Docker build configuration
├── docker-compose.yml   # Docker Compose setup
├── Makefile            # Build automation
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── LICENSE             # MIT License
├── README.md           # Project overview
├── CONTRIBUTING.md     # Contribution guidelines
├── TODOS.md            # Project TODOs
└── STRUCTURE.md        # This file
```

## Directory Descriptions

### `/cmd`
Contains the main applications for the project. The directory name matches the binary name.

- `cmd/server`: The main WhatsApp Business API server

### `/internal`
Private application code that shouldn't be imported by other projects.

- `api`: REST API routes, handlers, and middleware
- `database`: Database connection, migrations, and repository layer
- `models`: Domain models and database schemas
- `whatsapp`: WhatsApp Business Cloud API integration
- `mcp`: Model Context Protocol server for Claude AI
- `voice`: Voice calling functionality (SIP, WebRTC, transcription)

### `/pkg`
Public libraries that can be used by external applications.

- `config`: Configuration loading and validation
- `logger`: Structured logging utilities

### `/migrations`
Database migration files (SQL or Go-based).

### `/tests`
Test files organized by type:
- `integration`: Integration tests
- `unit`: Unit tests (most unit tests live alongside the code they test)

### `/storage`
Local file storage:
- `media`: WhatsApp media files (images, videos, documents)
- `recordings`: Voice call recordings (Phase 4)

### `/scripts`
Build scripts, deployment scripts, and utility scripts.

### `/deployments`
Deployment configurations:
- Docker Compose files
- Kubernetes manifests
- Prometheus/monitoring configs

### `/docs`
Project documentation.

## Configuration Files

- `.env.example`: Template for environment variables
- `.gitignore`: Files to ignore in Git
- `.dockerignore`: Files to ignore in Docker builds
- `Dockerfile`: Docker image build instructions
- `docker-compose.yml`: Multi-container Docker application
- `Makefile`: Build automation commands
- `go.mod` & `go.sum`: Go module dependencies

## Build & Run

See [Makefile](./Makefile) for available commands:

```bash
make build          # Build the application
make run           # Run locally
make test          # Run tests
make docker-build  # Build Docker image
make docker-up     # Start with Docker Compose
```

## Next Steps

1. Implement database layer (GORM + PostgreSQL)
2. Build REST API endpoints (Gin framework)
3. Integrate WhatsApp Business Cloud API
4. Implement webhook handlers
5. Build MCP server for Claude integration
6. Add voice calling features (Phase 4)

See [ROADMAP.md](./docs/ROADMAP.md) for the detailed development plan.
