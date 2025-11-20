package database

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// HealthStatus represents the database health status
type HealthStatus struct {
	Healthy       bool          `json:"healthy"`
	ResponseTime  time.Duration `json:"response_time_ms"`
	OpenConns     int           `json:"open_connections"`
	IdleConns     int           `json:"idle_connections"`
	WaitCount     int64         `json:"wait_count"`
	Error         string        `json:"error,omitempty"`
}

// HealthCheck performs a database health check with timeout
func HealthCheck(db *gorm.DB) HealthStatus {
	return HealthCheckWithContext(context.Background(), db, 5*time.Second)
}

// HealthCheckWithContext performs a database health check with custom context and timeout
func HealthCheckWithContext(ctx context.Context, db *gorm.DB, timeout time.Duration) HealthStatus {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	start := time.Now()
	status := HealthStatus{
		Healthy: false,
	}

	// Execute a simple query to check database connectivity
	var result int
	err := db.WithContext(ctx).Raw("SELECT 1").Scan(&result).Error

	status.ResponseTime = time.Since(start)

	if err != nil {
		status.Error = err.Error()
		return status
	}

	// Get connection statistics
	sqlDB, err := db.DB()
	if err != nil {
		status.Error = err.Error()
		return status
	}

	stats := sqlDB.Stats()
	status.OpenConns = stats.OpenConnections
	status.IdleConns = stats.Idle
	status.WaitCount = stats.WaitCount
	status.Healthy = true

	return status
}

// WaitForDatabase waits for the database to become available
func WaitForDatabase(db *gorm.DB, maxAttempts int, retryInterval time.Duration) error {
	for i := 0; i < maxAttempts; i++ {
		if err := PingDatabase(db); err == nil {
			return nil
		}

		if i < maxAttempts-1 {
			time.Sleep(retryInterval)
		}
	}

	return gorm.ErrInvalidDB
}
