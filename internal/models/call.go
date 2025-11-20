package models

import "time"

// Call represents a voice call record (Phase 4)
type Call struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	FromNumber     string    `json:"from_number" gorm:"index"`
	ToNumber       string    `json:"to_number" gorm:"index"`
	Direction      string    `json:"direction"` // inbound or outbound
	Status         string    `json:"status" gorm:"index"`
	Duration       int       `json:"duration"` // in seconds
	RecordingURL   string    `json:"recording_url,omitempty"`
	TranscriptID   string    `json:"transcript_id,omitempty"`
	StartedAt      time.Time `json:"started_at" gorm:"index"`
	EndedAt        *time.Time `json:"ended_at,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Transcript represents a call transcription (Phase 4)
type Transcript struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	CallID          string    `json:"call_id" gorm:"index"`
	Content         string    `json:"content" gorm:"type:text"`
	Language        string    `json:"language"`
	Provider        string    `json:"provider"` // whisper or deepgram
	Confidence      float64   `json:"confidence"`
	ProcessedAt     time.Time `json:"processed_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// TranscriptSegment represents a segment of a transcript with speaker info (Phase 4)
type TranscriptSegment struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	TranscriptID string    `json:"transcript_id" gorm:"index"`
	Speaker      string    `json:"speaker"`
	Content      string    `json:"content"`
	StartTime    float64   `json:"start_time"` // seconds from call start
	EndTime      float64   `json:"end_time"`
	Confidence   float64   `json:"confidence"`
	CreatedAt    time.Time `json:"created_at"`
}
