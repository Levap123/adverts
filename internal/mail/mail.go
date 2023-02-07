package mail

import (
	"github.com/Levap123/adverts/configs"
	"github.com/go-mail/mail"
)

type MailSender struct {
	from     string
	smtp     string
	smtpPort int
	password string
}

func NewMailSender(conf configs.EmailConf) *MailSender {
	return &MailSender{
		from:     conf.Email,
		smtp:     conf.Smtp,
		smtpPort: conf.SmtpPort,
		password: conf.Password,
	}
}

func (ms *MailSender) Send(to, message string) error {
	m := mail.NewMessage()
	m.SetHeader("From", ms.from)
	m.SetHeader("To", to)
	m.SetBody("text/plain", message)
	m.SetHeader("Subject", "adverts.com")
	d := mail.NewDialer(ms.smtp, ms.smtpPort, ms.from, ms.password)
	return d.DialAndSend(m)
}
