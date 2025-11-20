package validator

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var (
	// E.164 phone number format: +[country code][number]
	phoneRegex = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	// Email regex (basic)
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

// ValidatePhoneNumber validates a phone number in E.164 format
func ValidatePhoneNumber(phone string) error {
	if phone == "" {
		return errors.New("phone number is required")
	}
	if !phoneRegex.MatchString(phone) {
		return fmt.Errorf("invalid phone number format: must be E.164 format (e.g., +12345678901)")
	}
	return nil
}

// ValidateURL validates a URL
func ValidateURL(urlStr string) error {
	if urlStr == "" {
		return errors.New("URL is required")
	}
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	return nil
}

// ValidateMessageType validates a message type
func ValidateMessageType(msgType string) error {
	validTypes := []string{"text", "image", "video", "audio", "document", "location", "template"}
	for _, validType := range validTypes {
		if msgType == validType {
			return nil
		}
	}
	return fmt.Errorf("invalid message type: %s (must be one of: %s)", msgType, strings.Join(validTypes, ", "))
}

// ValidateStatus validates a message status
func ValidateStatus(status string) error {
	validStatuses := []string{"queued", "sent", "delivered", "read", "failed"}
	for _, validStatus := range validStatuses {
		if status == validStatus {
			return nil
		}
	}
	return fmt.Errorf("invalid status: %s (must be one of: %s)", status, strings.Join(validStatuses, ", "))
}

// ValidateLanguageCode validates an ISO 639-1 language code
func ValidateLanguageCode(code string) error {
	if code == "" {
		return errors.New("language code is required")
	}
	if len(code) != 2 && len(code) != 5 {
		return fmt.Errorf("invalid language code format: %s (must be ISO 639-1, e.g., 'en' or 'en_US')", code)
	}
	return nil
}

// ValidateEmail validates an email address
func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format: %s", email)
	}
	return nil
}

// ValidateNotEmpty validates that a string is not empty
func ValidateNotEmpty(value, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

// ValidateMaxLength validates maximum string length
func ValidateMaxLength(value, fieldName string, maxLength int) error {
	if len(value) > maxLength {
		return fmt.Errorf("%s exceeds maximum length of %d characters", fieldName, maxLength)
	}
	return nil
}

// ValidateMinLength validates minimum string length
func ValidateMinLength(value, fieldName string, minLength int) error {
	if len(value) < minLength {
		return fmt.Errorf("%s must be at least %d characters", fieldName, minLength)
	}
	return nil
}

// NormalizePhoneNumber normalizes a phone number to E.164 format
// This is a basic implementation - you may want to use a library like libphonenumber
func NormalizePhoneNumber(phone string) string {
	// Remove spaces, dashes, and parentheses
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	// Ensure it starts with +
	if !strings.HasPrefix(phone, "+") {
		phone = "+" + phone
	}

	return phone
}
