package mail

import (
	"net/smtp"

	"github.com/Levap123/adverts/configs"
)

type MailSender struct {
	auth       smtp.Auth
	from       string
	smtpServer string
}

func NewMailSender(conf configs.EmailConf) *MailSender {
	auth := smtp.PlainAuth("", conf.Email, conf.Password, conf.Smtp)
	return &MailSender{
		auth:       auth,
		from:       conf.Email,
		smtpServer: conf.Smtp + conf.SmtpPort,
	}
}

func (ms *MailSender) Send(to, message string) error {
	return smtp.SendMail(ms.smtpServer, ms.auth, ms.from, []string{to}, []byte(message))
}
