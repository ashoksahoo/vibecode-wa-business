package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all application configuration
type Config struct {
	// Server
	ServerPort int
	ServerHost string
	Environment string

	// Database
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// WhatsApp
	WhatsAppPhoneNumberID      string
	WhatsAppBusinessAccountID  string
	WhatsAppAccessToken        string
	WhatsAppWebhookVerifyToken string
	WhatsAppAPIVersion         string

	// Storage
	StorageType          string
	S3Bucket             string
	S3Region             string
	S3AccessKey          string
	S3SecretKey          string
	S3Endpoint           string
	MediaStoragePath     string
	RecordingsStoragePath string

	// MCP
	MCPEnabled bool
	MCPPort    int

	// Voice
	VoiceEnabled       bool
	SIPDomain          string
	SIPUsername        string
	SIPPassword        string
	WebRTCSTUNServers  string
	WebRTCTURNServers  string

	// Transcription
	TranscriptionProvider string
	WhisperModel         string
	DeepgramAPIKey       string

	// Logging
	LogLevel  string
	LogFormat string

	// Metrics
	MetricsEnabled bool
	MetricsPort    int

	// Rate Limiting
	RateLimitEnabled            bool
	RateLimitRequestsPerMinute int
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		ServerPort:                 getEnvAsInt("SERVER_PORT", 8080),
		ServerHost:                 getEnv("SERVER_HOST", "localhost"),
		Environment:                getEnv("ENV", "development"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "vibecoded_wa"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		WhatsAppPhoneNumberID:      getEnv("WHATSAPP_PHONE_NUMBER_ID", ""),
		WhatsAppBusinessAccountID:  getEnv("WHATSAPP_BUSINESS_ACCOUNT_ID", ""),
		WhatsAppAccessToken:        getEnv("WHATSAPP_ACCESS_TOKEN", ""),
		WhatsAppWebhookVerifyToken: getEnv("WHATSAPP_WEBHOOK_VERIFY_TOKEN", ""),
		WhatsAppAPIVersion:         getEnv("WHATSAPP_API_VERSION", "v18.0"),

		StorageType:           getEnv("STORAGE_TYPE", "local"),
		S3Bucket:              getEnv("S3_BUCKET", "vibecoded-wa-media"),
		S3Region:              getEnv("S3_REGION", "us-east-1"),
		S3AccessKey:           getEnv("S3_ACCESS_KEY", ""),
		S3SecretKey:           getEnv("S3_SECRET_KEY", ""),
		S3Endpoint:            getEnv("S3_ENDPOINT", ""),
		MediaStoragePath:      getEnv("MEDIA_STORAGE_PATH", "./storage/media"),
		RecordingsStoragePath: getEnv("RECORDINGS_STORAGE_PATH", "./storage/recordings"),

		MCPEnabled: getEnvAsBool("MCP_ENABLED", true),
		MCPPort:    getEnvAsInt("MCP_PORT", 3000),

		VoiceEnabled:      getEnvAsBool("VOICE_ENABLED", false),
		SIPDomain:         getEnv("SIP_DOMAIN", ""),
		SIPUsername:       getEnv("SIP_USERNAME", ""),
		SIPPassword:       getEnv("SIP_PASSWORD", ""),
		WebRTCSTUNServers: getEnv("WEBRTC_STUN_SERVERS", "stun:stun.l.google.com:19302"),
		WebRTCTURNServers: getEnv("WEBRTC_TURN_SERVERS", ""),

		TranscriptionProvider: getEnv("TRANSCRIPTION_PROVIDER", "whisper"),
		WhisperModel:         getEnv("WHISPER_MODEL", "base"),
		DeepgramAPIKey:       getEnv("DEEPGRAM_API_KEY", ""),

		LogLevel:  getEnv("LOG_LEVEL", "info"),
		LogFormat: getEnv("LOG_FORMAT", "json"),

		MetricsEnabled: getEnvAsBool("METRICS_ENABLED", true),
		MetricsPort:    getEnvAsInt("METRICS_PORT", 9090),

		RateLimitEnabled:           getEnvAsBool("RATE_LIMIT_ENABLED", true),
		RateLimitRequestsPerMinute: getEnvAsInt("RATE_LIMIT_REQUESTS_PER_MINUTE", 60),
	}

	// Validate required fields
	if cfg.Environment == "production" {
		if cfg.WhatsAppPhoneNumberID == "" {
			return nil, fmt.Errorf("WHATSAPP_PHONE_NUMBER_ID is required in production")
		}
		if cfg.WhatsAppAccessToken == "" {
			return nil, fmt.Errorf("WHATSAPP_ACCESS_TOKEN is required in production")
		}
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}
