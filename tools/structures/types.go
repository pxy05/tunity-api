package structures

import (
	"time"
)

type User struct {
	ID        *string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	Username  *string    `json:"username,omitempty"`
	Email     *string    `json:"email,omitempty"`
	FirstName *string    `json:"first_name,omitempty"`
	LastName  *string    `json:"last_name,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Interview struct {
	ID        *string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	AppliID   *string    `json:"appli_id,omitempty"`
	Round     *int16     `json:"round,omitempty"`
	Date      *time.Time `json:"date,omitempty"`
	Type      *string    `json:"type,omitempty"`
	Result    *string    `json:"result,omitempty"`
}

type Application struct {
	// Status can be: applied, rejected, interview, offer, offer_accepted, offer_rejected
	ID                    *string     `json:"id,omitempty"`
	UserID                *string    `json:"user_id"`
	CreatedAt             time.Time  `json:"created_at,omitempty"`
	UpdatedAt             *time.Time `json:"updated_at,omitempty"`
	Company               *string    `json:"company"`
	Status                *string    `json:"status,omitempty"`
	AppliTitle            *string    `json:"appli_title,omitempty"`
	AppliLocation         *string    `json:"appli_location,omitempty"`
	AppliNotes            *string    `json:"appli_notes,omitempty"`
	AppliURL              *string    `json:"appli_url,omitempty"`
	AppliDeadline         *time.Time `json:"appli_deadline,omitempty"`
	AppliRejected         *time.Time `json:"appli_rejected,omitempty"`
	InterviewFailedDate   *time.Time `json:"interview_failed_date,omitempty"`
	InterviewFailedReason *string    `json:"interview_failed_reason,omitempty"`
}