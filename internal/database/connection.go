package database

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PoolConfig holds database connection pool settings
type PoolConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// DefaultPoolConfig returns the default connection pool configuration
func DefaultPoolConfig() PoolConfig {
	return PoolConfig{
		MaxIdleConns:    5,
		MaxOpenConns:    25,
		ConnMaxLifetime: time.Hour,
	}
}

// NewConnection creates a new database connection
func NewConnection(driver, dsn string, logLevel logger.LogLevel) (*gorm.DB, error) {
	var dialector gorm.Dialector

	// Select the appropriate driver
	switch driver {
	case "sqlite":
		dialector = sqlite.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s (supported: sqlite, postgres)", driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get underlying SQL database for connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set default connection pool settings
	SetConnectionPool(sqlDB, DefaultPoolConfig())

	// Ping to verify connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// SetConnectionPool configures the database connection pool
func SetConnectionPool(sqlDB *sql.DB, config PoolConfig) {
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
}

// CloseConnection closes the database connection
func CloseConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	return nil
}

// PingDatabase checks if the database is reachable
func PingDatabase(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	return nil
}

// GetConnectionStats returns database connection statistics
func GetConnectionStats(db *gorm.DB) (sql.DBStats, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return sql.DBStats{}, fmt.Errorf("failed to get database instance: %w", err)
	}

	return sqlDB.Stats(), nil
}
