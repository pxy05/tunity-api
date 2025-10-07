package structures

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Interview struct {
	gorm.Model
	ApplicationID uint   `json:"application_id" gorm:"not null;foreignKey:ApplicationID;references:ID"`
	Round         int    `json:"round"`
	Date          *time.Time `json:"date"`
	Type          *string `json:"type"`
	Result        *string `json:"result"`
}

type Application struct {
	// Status can be: applied, rejected, interview, offer, offer_accepted, offer_rejected
	gorm.Model
	UserID   uint   `json:"user_id" gorm:"not null;foreignKey:UserID;references:ID"`
	Company string `json:"company"`
	Status string `json:"status"`
	Job_Title string `json:"job_title"`
	Job_Location string `json:"job_location"`
	Job_Notes *string `json:"job_notes"`

	Application_URL *string `json:"application_url"`
	Application_Deadline *time.Time `json:"application_deadline"`
	Application_Applied_Date *time.Time `json:"application_applied_date"`
	Application_Rejected_Date *time.Time `json:"application_rejected_date"`
	Interview_Failed_Date *time.Time `json:"interview_failed_date"`
	Interview_Failed_Reason *string `json:"interview_failed_reason" gorm:"type:text"`
}