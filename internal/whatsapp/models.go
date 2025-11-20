package whatsapp

import "time"

// MediaType represents the type of media
type MediaType string

const (
	MediaTypeImage    MediaType = "image"
	MediaTypeDocument MediaType = "document"
	MediaTypeAudio    MediaType = "audio"
	MediaTypeVideo    MediaType = "video"
)

// MessageRequest represents a request to send a message
type MessageRequest struct {
	Phone            string            `json:"phone"`
	Type             string            `json:"type"`
	Content          string            `json:"content,omitempty"`
	MediaURL         string            `json:"media_url,omitempty"`
	Caption          string            `json:"caption,omitempty"`
	Filename         string            `json:"filename,omitempty"`
	TemplateName     string            `json:"template_name,omitempty"`
	TemplateLanguage string            `json:"template_language,omitempty"`
	Parameters       []string          `json:"parameters,omitempty"`
}

// MessageResponse represents WhatsApp API response
type MessageResponse struct {
	MessagingProduct string `json:"messaging_product"`
	Contacts         []struct {
		Input string `json:"input"`
		WaID  string `json:"wa_id"`
	} `json:"contacts"`
	Messages []struct {
		ID string `json:"id"`
	} `json:"messages"`
}

// MessageStatus represents the status of a message
type MessageStatus struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// WebhookPayload represents incoming webhook data
type WebhookPayload struct {
	Object string `json:"object"`
	Entry  []struct {
		ID      string `json:"id"`
		Changes []struct {
			Value struct {
				MessagingProduct string          `json:"messaging_product"`
				Metadata         MetadataValue   `json:"metadata"`
				Contacts         []ContactValue  `json:"contacts,omitempty"`
				Messages         []MessageValue  `json:"messages,omitempty"`
				Statuses         []StatusValue   `json:"statuses,omitempty"`
			} `json:"value"`
			Field string `json:"field"`
		} `json:"changes"`
	} `json:"entry"`
}

// MetadataValue represents webhook metadata
type MetadataValue struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

// ContactValue represents contact information in webhook
type ContactValue struct {
	Profile struct {
		Name string `json:"name"`
	} `json:"profile"`
	WaID string `json:"wa_id"`
}

// MessageValue represents a message in webhook
type MessageValue struct {
	From      string `json:"from"`
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Text      *struct {
		Body string `json:"body"`
	} `json:"text,omitempty"`
	Image *struct {
		Caption  string `json:"caption,omitempty"`
		MimeType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
		ID       string `json:"id"`
	} `json:"image,omitempty"`
	Document *struct {
		Caption  string `json:"caption,omitempty"`
		Filename string `json:"filename,omitempty"`
		MimeType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
		ID       string `json:"id"`
	} `json:"document,omitempty"`
	Audio *struct {
		MimeType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
		ID       string `json:"id"`
		Voice    bool   `json:"voice"`
	} `json:"audio,omitempty"`
	Video *struct {
		Caption  string `json:"caption,omitempty"`
		MimeType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
		ID       string `json:"id"`
	} `json:"video,omitempty"`
}

// StatusValue represents a status update in webhook
type StatusValue struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
	RecipientID string `json:"recipient_id"`
	Errors      []struct {
		Code    int    `json:"code"`
		Title   string `json:"title"`
		Message string `json:"message,omitempty"`
	} `json:"errors,omitempty"`
}

// MessageEvent represents a parsed incoming message event
type MessageEvent struct {
	MessageID   string
	From        string
	Timestamp   time.Time
	Type        string
	Content     string
	MediaID     string
	MediaURL    string
	MimeType    string
	Caption     string
	Filename    string
	ContactName string
}

// StatusEvent represents a parsed status update event
type StatusEvent struct {
	MessageID   string
	Status      string
	Timestamp   time.Time
	RecipientID string
	ErrorCode   int
	ErrorTitle  string
	ErrorMsg    string
}

// ErrorResponse represents an error from WhatsApp API
type ErrorResponse struct {
	Error struct {
		Message      string `json:"message"`
		Type         string `json:"type"`
		Code         int    `json:"code"`
		ErrorData    string `json:"error_data,omitempty"`
		ErrorSubcode int    `json:"error_subcode,omitempty"`
		FBTraceID    string `json:"fbtrace_id,omitempty"`
	} `json:"error"`
}
