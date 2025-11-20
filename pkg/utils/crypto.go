package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GenerateAPIKey generates a secure random API key
func GenerateAPIKey() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// HashAPIKey hashes an API key using bcrypt
func HashAPIKey(key string) (string, error) {
	hashedKey, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash API key: %w", err)
	}
	return string(hashedKey), nil
}

// CompareAPIKey compares a hashed key with a plain key
func CompareAPIKey(hashedKey, key string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedKey), []byte(key))
	return err == nil
}

// GenerateID generates a prefixed ID (e.g., "msg_abc123")
func GenerateID(prefix string) string {
	id := uuid.New().String()
	if prefix != "" {
		return fmt.Sprintf("%s_%s", prefix, id)
	}
	return id
}

// GetKeyPrefix extracts the first 8 characters of a key for identification
func GetKeyPrefix(key string) string {
	if len(key) < 8 {
		return key
	}
	return key[:8]
}

// ComputeHMAC computes HMAC-SHA256 for a message
func ComputeHMAC(message, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(message)
	return "sha256=" + hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC verifies an HMAC signature (constant-time comparison)
func VerifyHMAC(message, secret []byte, signature string) bool {
	expectedMAC := ComputeHMAC(message, secret)
	return hmac.Equal([]byte(signature), []byte(expectedMAC))
}

// GenerateRandomString generates a random string of specified length
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate random string: %w", err)
	}
	return hex.EncodeToString(b)[:length], nil
}
