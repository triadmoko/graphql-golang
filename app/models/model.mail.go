package models

import "time"

type (
	FormSendEmail struct {
		Subject string
		Text    string
		To      []string
	}
	EmailVerification struct {
		ID        string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
		UserID    string
		Email     string
		Code      int
		Expired   int64
	}
)
