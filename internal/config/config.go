package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	WhatsApp WhatsAppConfig
	Security SecurityConfig
	Logging  LoggingConfig
	Metrics  MetricsConfig
	Storage  StorageConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port            int
	Host            string
	Environment     string // development, staging, production
	BaseURL         string
	ShutdownTimeout time.Duration
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Driver          string // sqlite or postgres
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	SQLitePath      string // Path to SQLite database file
}

// WhatsAppConfig holds WhatsApp API configuration
type WhatsAppConfig struct {
	APIToken            string
	PhoneNumberID       string
	BusinessAccountID   string
	WebhookVerifyToken  string
	WebhookSecret       string
	APIBaseURL          string
	APIVersion          string
}

// SecurityConfig holds security configuration
type SecurityConfig struct {
	APIKeySalt    string
	SessionSecret string
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level      string // debug, info, warn, error
	Format     string // json, console
	OutputPath string
}

// MetricsConfig holds metrics configuration
type MetricsConfig struct {
	Enabled bool
	Port    int
}

// StorageConfig holds storage configuration
type StorageConfig struct {
	Type              string // local, s3, minio
	S3Bucket          string
	S3Region          string
	S3AccessKey       string
	S3SecretKey       string
	S3Endpoint        string
	MediaPath         string
	RecordingsPath    string
}

// LoadConfig loads configuration from environment variables and .env file
func LoadConfig() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./")

	// Read .env file if it exists
	_ = viper.ReadInConfig()

	// Bind environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config := &Config{
		Server: ServerConfig{
			Port:            viper.GetInt("SERVER_PORT"),
			Host:            viper.GetString("SERVER_HOST"),
			Environment:     viper.GetString("ENV"),
			BaseURL:         viper.GetString("SERVER_BASE_URL"),
			ShutdownTimeout: viper.GetDuration("SERVER_SHUTDOWN_TIMEOUT"),
		},
		Database: DatabaseConfig{
			Driver:          viper.GetString("DB_DRIVER"),
			Host:            viper.GetString("DB_HOST"),
			Port:            viper.GetInt("DB_PORT"),
			User:            viper.GetString("DB_USER"),
			Password:        viper.GetString("DB_PASSWORD"),
			Name:            viper.GetString("DB_NAME"),
			SSLMode:         viper.GetString("DB_SSL_MODE"),
			SQLitePath:      viper.GetString("DB_SQLITE_PATH"),
			MaxOpenConns:    viper.GetInt("DB_MAX_OPEN_CONNS"),
			MaxIdleConns:    viper.GetInt("DB_MAX_IDLE_CONNS"),
			ConnMaxLifetime: viper.GetDuration("DB_CONN_MAX_LIFETIME"),
		},
		WhatsApp: WhatsAppConfig{
			APIToken:            viper.GetString("WHATSAPP_ACCESS_TOKEN"),
			PhoneNumberID:       viper.GetString("WHATSAPP_PHONE_NUMBER_ID"),
			BusinessAccountID:   viper.GetString("WHATSAPP_BUSINESS_ACCOUNT_ID"),
			WebhookVerifyToken:  viper.GetString("WHATSAPP_WEBHOOK_VERIFY_TOKEN"),
			WebhookSecret:       viper.GetString("WHATSAPP_WEBHOOK_SECRET"),
			APIBaseURL:          viper.GetString("WHATSAPP_API_BASE_URL"),
			APIVersion:          viper.GetString("WHATSAPP_API_VERSION"),
		},
		Security: SecurityConfig{
			APIKeySalt:    viper.GetString("API_KEY_SALT"),
			SessionSecret: viper.GetString("SESSION_SECRET"),
		},
		Logging: LoggingConfig{
			Level:      viper.GetString("LOG_LEVEL"),
			Format:     viper.GetString("LOG_FORMAT"),
			OutputPath: viper.GetString("LOG_OUTPUT_PATH"),
		},
		Metrics: MetricsConfig{
			Enabled: viper.GetBool("METRICS_ENABLED"),
			Port:    viper.GetInt("METRICS_PORT"),
		},
		Storage: StorageConfig{
			Type:           viper.GetString("STORAGE_TYPE"),
			S3Bucket:       viper.GetString("S3_BUCKET"),
			S3Region:       viper.GetString("S3_REGION"),
			S3AccessKey:    viper.GetString("S3_ACCESS_KEY"),
			S3SecretKey:    viper.GetString("S3_SECRET_KEY"),
			S3Endpoint:     viper.GetString("S3_ENDPOINT"),
			MediaPath:      viper.GetString("MEDIA_STORAGE_PATH"),
			RecordingsPath: viper.GetString("RECORDINGS_STORAGE_PATH"),
		},
	}

	// Set defaults
	setDefaults(config)

	// Validate configuration
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return config, nil
}

// setDefaults sets default values for configuration
func setDefaults(config *Config) {
	if config.Server.Port == 0 {
		config.Server.Port = 8080
	}
	if config.Server.Host == "" {
		config.Server.Host = "localhost"
	}
	if config.Server.Environment == "" {
		config.Server.Environment = "development"
	}
	if config.Server.ShutdownTimeout == 0 {
		config.Server.ShutdownTimeout = 30 * time.Second
	}

	if config.Database.Driver == "" {
		config.Database.Driver = "sqlite" // Default to SQLite for development
	}
	if config.Database.SQLitePath == "" {
		config.Database.SQLitePath = "./vibecoded.db"
	}
	if config.Database.Port == 0 {
		config.Database.Port = 5432
	}
	if config.Database.SSLMode == "" {
		config.Database.SSLMode = "disable"
	}
	if config.Database.MaxOpenConns == 0 {
		config.Database.MaxOpenConns = 25
	}
	if config.Database.MaxIdleConns == 0 {
		config.Database.MaxIdleConns = 5
	}
	if config.Database.ConnMaxLifetime == 0 {
		config.Database.ConnMaxLifetime = time.Hour
	}

	if config.WhatsApp.APIBaseURL == "" {
		config.WhatsApp.APIBaseURL = "https://graph.facebook.com"
	}
	if config.WhatsApp.APIVersion == "" {
		config.WhatsApp.APIVersion = "v18.0"
	}

	if config.Logging.Level == "" {
		config.Logging.Level = "info"
	}
	if config.Logging.Format == "" {
		config.Logging.Format = "json"
	}

	if config.Metrics.Port == 0 {
		config.Metrics.Port = 9090
	}

	if config.Storage.Type == "" {
		config.Storage.Type = "local"
	}
	if config.Storage.MediaPath == "" {
		config.Storage.MediaPath = "./storage/media"
	}
	if config.Storage.RecordingsPath == "" {
		config.Storage.RecordingsPath = "./storage/recordings"
	}
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Server.Port < 0 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", c.Server.Port)
	}

	if c.Server.Environment == "production" {
		if c.WhatsApp.APIToken == "" {
			return fmt.Errorf("WHATSAPP_ACCESS_TOKEN is required in production")
		}
		if c.WhatsApp.PhoneNumberID == "" {
			return fmt.Errorf("WHATSAPP_PHONE_NUMBER_ID is required in production")
		}
		if c.Database.Password == "" {
			return fmt.Errorf("DB_PASSWORD is required in production")
		}
	}

	return nil
}

// GetDatabaseDSN returns the database connection string
func (c *Config) GetDatabaseDSN() string {
	if c.Database.Driver == "sqlite" {
		return c.Database.SQLitePath
	}

	// PostgreSQL DSN
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

// GetDatabaseDriver returns the database driver name
func (c *Config) GetDatabaseDriver() string {
	return c.Database.Driver
}

// IsDevelopment returns true if running in development mode
func (c *Config) IsDevelopment() bool {
	return c.Server.Environment == "development"
}

// IsProduction returns true if running in production mode
func (c *Config) IsProduction() bool {
	return c.Server.Environment == "production"
}
