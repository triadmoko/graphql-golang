package models

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
)

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

func FormatterRequestVerificatonEmail(email, userID string) EmailVerification {
	min := 000000
	max := 999999
	code := rand.Intn(max-min) + min
	response := EmailVerification{
		ID:        helpers.UUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
		UserID:    userID,
		Email:     email,
		Code:      code,
		Expired:   time.Now().UTC().Add(1 * time.Minute).Unix(),
	}
	return response
}

func FormatterRequestSendEmail(request EmailVerification) FormSendEmail {

	response := FormSendEmail{
		Subject: "",
		Text:    strconv.Itoa(request.Code),
		To:      []string{request.Email},
	}
	return response
}
