package mail

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
)

func FormatterRequestVerificatonEmail(email, userID string) models.EmailVerification {
	min := 000000
	max := 999999
	code := rand.Intn(max-min) + min
	response := models.EmailVerification{
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

func FormatterRequestSendEmail(request models.EmailVerification) models.FormSendEmail {

	response := models.FormSendEmail{
		Subject: "",
		Text:    strconv.Itoa(request.Code),
		To:      []string{request.Email},
	}
	return response
}
