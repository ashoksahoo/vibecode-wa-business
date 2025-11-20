# Dependencies

## External Services

| Service | Purpose | Required | Notes |
|---------|---------|----------|-------|
| WhatsApp Cloud API | Message sending/receiving | Yes | Free tier available for development |
| PostgreSQL | Primary database | Yes | Version 15+ recommended |
| Docker | Containerization | Yes | For deployment |
| Meta Developer Account | WhatsApp API access | Yes | Free to create |

## Go Libraries

### Core Web Framework
```go
github.com/gin-gonic/gin    // HTTP web framework
```

### Database
```go
gorm.io/gorm                 // ORM
gorm.io/driver/postgres      // PostgreSQL driver
```

### HTTP Client
```go
github.com/go-resty/resty/v2 // HTTP client for WhatsApp API calls
```

### Configuration & Environment
```go
github.com/spf13/viper       // Configuration management
github.com/joho/godotenv     // .env file support
```

### Logging
```go
go.uber.org/zap              // Structured logging
```

### Validation
```go
github.com/go-playground/validator/v10  // Request validation
```

### Utilities
```go
github.com/google/uuid       // UUID generation
```

### Testing
```go
github.com/stretchr/testify  // Test assertions
```

### Metrics (Optional)
```go
github.com/prometheus/client_golang  // Prometheus metrics
```

## Development Tools

### Required
- **Go**: Version 1.21 or higher
- **Docker**: Latest stable version
- **Docker Compose**: For local development
- **PostgreSQL**: Version 15+ (via Docker)
- **Make**: For build automation
- **Git**: Version control

### Recommended
- **curl**: For API testing
- **jq**: For JSON parsing in scripts
- **pgcli**: PostgreSQL command-line interface
- **air**: For hot-reloading during development

## System Requirements

### Minimum
- 2 CPU cores
- 2GB RAM
- 10GB disk space
- Linux/macOS/Windows with WSL2

### Recommended
- 4 CPU cores
- 4GB RAM
- 20GB disk space
- Linux or macOS for development

## API Requirements

### WhatsApp Business Platform
- Business verification (for production)
- Phone number verification
- WhatsApp Business API access token
- Webhook URL (HTTPS required for production)

## Network Requirements

### Inbound
- Port 8080 (REST API)
- Port for webhooks (configurable)

### Outbound
- HTTPS to graph.facebook.com (WhatsApp API)
- PostgreSQL port (default 5432)

## TODO for Claude Code

- [ ] Create Go module with all dependencies
- [ ] Generate Docker Compose configuration with PostgreSQL
- [ ] Create Makefile for common development tasks
- [ ] Generate dependency installation script
