package injector

import (
	"github.com/triadmoko/grahpql-golang/config"
	mailRepository "github.com/triadmoko/grahpql-golang/repository/mail"
	mailService "github.com/triadmoko/grahpql-golang/service/mail"
)

func NewInjectorMail(conf *config.Config) mailService.MailServices {
	repository := mailRepository.NewMailRepositorys(conf)
	service := mailService.NewMailServices(conf.Logger, repository)
	return service
}
