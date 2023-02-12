package mail

import (
	"context"

	"github.com/triadmoko/grahpql-golang/models"
)

func (s *mail_service) EmailVerification(ctx context.Context, email, userID string) error {
	request := models.FormatterRequestVerificatonEmail(email, userID)
	result, err := s.mailRepository.Create(ctx, request)
	if err != nil {
		s.loggger.Error("Error create email on verifycation email")
		return err
	}
	requestSendEmail := models.FormatterRequestSendEmail(*result)
	err = s.SendEmail(requestSendEmail)
	if err != nil {
		s.loggger.Error(err.Error())
		return err
	}
	return nil
}
