# Technical Requirements
# Vibecoded WA Client

**Last Updated:** November 21, 2025

---

## Core Requirements

### API Design
- ✅ RESTful API with proper HTTP methods (GET, POST, PUT, PATCH, DELETE)
- ✅ OpenAPI/Swagger documentation
- ✅ Versioned API endpoints (/api/v1)
- ✅ Consistent response format
- ✅ Proper HTTP status codes
- ✅ Content-Type negotiation

**TODO for Claude Code:** OpenAPI specification generator

---

### Data Persistence
- ✅ PostgreSQL database integration
- ✅ Database schema migrations
- ✅ Transaction support
- ✅ Connection pooling
- ✅ Query optimization
- ✅ Database indexes

**TODO for Claude Code:** Database connection pool configuration

---

### Configuration Management
- ✅ Environment-based configuration
- ✅ Support for .env files
- ✅ Configuration validation on startup
- ✅ Secrets management (no hardcoded values)
- ✅ Configuration documentation

**TODO for Claude Code:** Configuration validation module

---

### Logging & Monitoring
- ✅ Structured logging (JSON format)
- ✅ Log levels (DEBUG, INFO, WARN, ERROR)
- ✅ Request/response logging
- ✅ Correlation IDs for request tracing
- ✅ Error stack traces

**TODO for Claude Code:** Structured logging setup with zap

---

### Input Validation
- ✅ Request body validation
- ✅ Query parameter validation
- ✅ Phone number format validation
- ✅ URL validation for media
- ✅ File size validation
- ✅ Input sanitization

**TODO for Claude Code:** Validation middleware

---

### Error Handling
- ✅ Consistent error response format
- ✅ Error codes and messages
- ✅ Graceful error recovery
- ✅ Error logging
- ✅ User-friendly error messages

---

### Testing
- ✅ Unit tests for core logic (>70% coverage)
- ✅ Integration tests for API endpoints
- ✅ Test fixtures and mocks
- ✅ Test database setup/teardown
- ✅ CI/CD pipeline integration

**TODO for Claude Code:** Test suite infrastructure

---

### Deployment
- ✅ Docker containerization
- ✅ Multi-stage Docker build
- ✅ Docker Compose for local development
- ✅ Health check endpoint
- ✅ Graceful shutdown
- ✅ Environment variable configuration

**TODO for Claude Code:** Dockerfile and docker-compose.yml

---

### Authentication & Authorization
- 📋 JWT token-based authentication
- 📋 API key authentication for programmatic access
- 📋 Role-based access control (RBAC)
- 📋 Permission system
- 📋 Token refresh mechanism
- 📋 Session management

**TODO for Claude Code:** JWT middleware
**TODO for Claude Code:** RBAC middleware

---

### Rate Limiting
- 📋 Request rate limiting per API key
- 📋 Per-user rate limits
- 📋 Rate limit headers (X-RateLimit-*)
- 📋 Rate limit exceeded responses
- 📋 Configurable rate limits

**TODO for Claude Code:** Rate limiting middleware

---

### Request/Response Handling
- 📋 Request ID generation and tracking
- 📋 CORS configuration
- 📋 Request timeout handling
- 📋 Response compression
- 📋 Pagination helpers

**TODO for Claude Code:** CORS and compression middleware

---

### Metrics & Observability
- 📋 Prometheus metrics endpoint
- 📋 Request duration metrics
- 📋 Error rate metrics
- 📋 Database query metrics
- 📋 Custom business metrics

**TODO for Claude Code:** Prometheus metrics integration

---

### Database Management
- 📋 Automated migrations
- 📋 Migration rollback support
- 📋 Seed data for development
- 📋 Database backup scripts
- 📋 Connection retry logic

**TODO for Claude Code:** Migration system

---

### Configuration
- 📋 Configuration hot-reload (where applicable)
- 📋 Feature flags
- 📋 Environment-specific configs
- 📋 Configuration API endpoint

---

### End-to-End Testing
- 📋 E2E test suite
- 📋 Test data generation
- 📋 API contract testing
- 📋 Performance testing

**TODO for Claude Code:** E2E test suite

---

### Advanced Features
- 💭 Admin API for system management
- 💭 CLI tool for common operations
- 💭 Multiple database support (MySQL, SQLite)
- 💭 Caching layer (Redis)
- 💭 Message queue (RabbitMQ/Redis)
- 💭 Distributed tracing (OpenTelemetry)
- 💭 Load testing suite (k6)

---

### Developer Experience
- 💭 API client SDKs (Go, Python, Node.js)
- 💭 Postman collection
- 💭 GraphQL API (alternative to REST)
- 💭 WebSocket support for real-time updates
- 💭 API playground/sandbox

---

## Performance Requirements

### Response Times
- API response time: **< 200ms** (p95)
- Webhook processing: **< 1s**
- Database queries: **< 50ms** (p95)
- Health check: **< 50ms**
- Message send: **< 500ms** (including WhatsApp API call)

### Throughput
- Handle **1000+ messages/hour**
- Support **10,000+ contacts**
- Store **1M+ messages**
- Process **100+ concurrent webhook requests**
- Support **100+ concurrent API connections**

### Scalability
- Horizontal scaling capability (stateless design)
- Database connection pooling (min: 5, max: 20)
- Efficient memory usage (< 512MB for typical workload)
- No memory leaks
- Graceful degradation under load

**TODO for Claude Code:** Performance benchmarking suite

---

## Security Requirements

### Transport Security
- HTTPS only for webhooks
- TLS 1.2+ required
- Certificate validation
- Secure headers (HSTS, CSP, etc.)

### Authentication & Authorization
- WhatsApp webhook signature verification
- API key authentication with Bearer tokens
- JWT token expiration (15 minutes)
- Refresh token rotation
- Password hashing (bcrypt, cost factor 12)
- API key hashing before storage

### Input Security
- Input validation and sanitization
- SQL injection prevention (parameterized queries)
- XSS prevention
- CSRF protection
- File upload validation
- Phone number format validation

### Data Security
- Environment variable secrets (no hardcoding)
- Database credentials encrypted at rest
- Sensitive data redaction in logs
- Audit logging for critical operations
- Data encryption in transit

### Rate Limiting & DDoS Protection
- Rate limiting per API key (1000 req/hour default)
- Rate limiting per IP address
- Request size limits (10MB max)
- Connection limits
- Timeout configuration

**TODO for Claude Code:** Security middleware suite
**TODO for Claude Code:** Webhook signature verification

---

## Reliability Requirements

### Availability
- System uptime: **> 99.9%**
- Planned maintenance windows
- Zero-downtime deployments
- Automatic failover capability

### Data Integrity
- ACID transaction support
- Data validation before persistence
- Foreign key constraints
- Unique constraints
- Audit trail for critical operations

### Error Recovery
- Graceful error handling
- Retry logic with exponential backoff
- Circuit breaker pattern
- Dead letter queue for failed messages
- Database connection retry

### Message Delivery
- Message delivery success: **> 99%**
- Idempotent message processing
- Duplicate message detection
- Message status tracking
- Delivery receipts

**TODO for Claude Code:** Circuit breaker implementation
**TODO for Claude Code:** Message retry queue

---

## Operational Requirements

### Monitoring
- Health check endpoint with detailed status
- Readiness and liveness probes
- System resource monitoring
- Application metrics
- Alert thresholds

### Logging
- Structured JSON logs
- Log aggregation ready
- Log rotation
- Configurable log levels
- Request/response logging (with PII redaction)

### Maintenance
- Database backup procedures
- Data migration tools
- Configuration management
- Version management
- Rollback capability

**TODO for Claude Code:** Health check implementation
**TODO for Claude Code:** Log rotation configuration

---

## Technology Stack

### Core
- **Language:** Go 1.21+
- **Web Framework:** Gin
- **ORM:** GORM
- **Database:** PostgreSQL 15+
- **Container:** Docker
- **Orchestration:** Docker Compose (local), Kubernetes (optional)

### Libraries
```go
// Web framework
github.com/gin-gonic/gin

// Database
gorm.io/gorm
gorm.io/driver/postgres

// HTTP client
github.com/go-resty/resty/v2

// Configuration
github.com/spf13/viper
github.com/joho/godotenv

// Logging
go.uber.org/zap

// Validation
github.com/go-playground/validator/v10

// Authentication
github.com/golang-jwt/jwt/v5
golang.org/x/crypto/bcrypt

// Testing
github.com/stretchr/testify

// UUID
github.com/google/uuid

// Metrics
github.com/prometheus/client_golang
```

**TODO for Claude Code:** Dependencies management (go.mod)

---

## Compliance & Standards

### API Standards
- RESTful API design principles
- OpenAPI 3.0 specification
- Semantic versioning
- Consistent naming conventions

### Code Standards
- Go standard project layout
- Effective Go guidelines
- Error handling best practices
- Code documentation (godoc)
- Code review process

### Data Standards
- ISO 8601 for timestamps
- E.164 for phone numbers
- UTF-8 encoding
- JSON response format
- Pagination standards (limit/offset)

**TODO for Claude Code:** Code linting configuration (golangci-lint)

---

## Development Requirements

### Version Control
- Git for source control
- Semantic versioning (semver)
- Conventional commits
- Branch protection rules
- Tag releases

### Documentation
- README with setup instructions
- API documentation (OpenAPI/Swagger)
- Architecture documentation
- Deployment guide
- Contributing guidelines

### Testing
- Test-driven development (TDD) encouraged
- Minimum 70% code coverage
- Integration test suite
- E2E test suite
- Performance test suite

**TODO for Claude Code:** Test coverage reporting

---

## Browser/Client Support

### Supported Clients
- cURL
- Postman
- HTTPie
- Programming language HTTP clients
- Claude Desktop (MCP)

### API Versioning
- URL-based versioning (/api/v1)
- Backward compatibility for minor versions
- Deprecation notices (6 months minimum)
- Migration guides for breaking changes

---

## Dependencies & Prerequisites

### Required
- Go 1.21 or higher
- PostgreSQL 15 or higher
- Docker 20.10 or higher
- Docker Compose 2.0 or higher
- WhatsApp Business API credentials
- Meta Developer Account

### Optional
- Redis (for caching and queuing)
- Prometheus (for metrics)
- Grafana (for dashboards)
- Make (for build automation)

---

## Constraints & Limitations

### WhatsApp API Limits
- Rate limits imposed by WhatsApp Cloud API
- Message template approval required
- Media file size limits (5MB for images, 16MB for documents)
- Phone number verification required

### System Limits
- Maximum message size: 4096 characters
- Maximum request body size: 10MB
- Maximum concurrent connections: 100
- Database connection pool: 20 max connections

### Feature Limitations
- No support for personal WhatsApp accounts
- Single WhatsApp Business number per instance
- No end-to-end encryption beyond WhatsApp's own
- No message editing or deletion after sending

---

## Quality Attributes

### Maintainability
- Clear code structure
- Comprehensive documentation
- Modular design
- Dependency injection
- Interface-based design

### Testability
- Unit test coverage > 70%
- Integration test coverage for all endpoints
- Mock-friendly architecture
- Test fixtures and helpers
- CI/CD integration

### Extensibility
- Plugin architecture (future)
- Webhook forwarding support
- Custom middleware support
- Event system for extensions

**TODO for Claude Code:** Plugin system architecture

---

## Performance Targets

| Requirement | Target |
|-------------|--------|
| API Response Time (p95) | < 200ms |
| Uptime | > 99.9% |
| Code Coverage | > 70% |
| Message Delivery Rate | > 99% |
| Concurrent Connections | > 100 |
| Messages per Hour | > 1000 |
| Database Query Time (p95) | < 50ms |
| Memory Usage (typical) | < 512MB |
