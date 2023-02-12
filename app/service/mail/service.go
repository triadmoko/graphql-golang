package mail

import (
	"context"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/models"
	"github.com/triadmoko/grahpql-golang/repository/mail"
	"gopkg.in/gomail.v2"
)

type MailServices interface {
	EmailVerification(ctx context.Context, email, userID string) error
	SendEmail(form models.FormSendEmail) error
}
type mail_service struct {
	loggger        *logrus.Logger
	dialer         *gomail.Dialer
	mailRepository mail.MailRepository
}

func NewMailServices(loggger *logrus.Logger, mailRepository mail.MailRepository) *mail_service {
	CONFIG_SMTP_PORT, _ := strconv.Atoi(helpers.LoadEnv("CONFIG_SMTP_PORT"))

	dialer := gomail.NewDialer(
		helpers.LoadEnv("CONFIG_SMTP_HOST"),
		CONFIG_SMTP_PORT,
		helpers.LoadEnv("CONFIG_AUTH_EMAIL"),
		helpers.LoadEnv("CONFIG_AUTH_PASSWORD"),
	)

	return &mail_service{loggger: loggger, mailRepository: mailRepository, dialer: dialer}
}
