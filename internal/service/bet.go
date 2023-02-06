package service

import (
	"github.com/Levap123/adverts/configs"
	"github.com/Levap123/adverts/internal/mail"
	"github.com/Levap123/adverts/internal/repository"
	"github.com/spf13/viper"
)

type Bet struct {
	repo       repository.BetRepo
	mailSender *mail.MailSender
}

func NewBet(repo repository.BetRepo) *Bet {
	maiSender := mail.NewMailSender(configs.EmailConf{
		Email:    viper.GetString("email"),
		Password: viper.GetString("password"),
		Smtp:     viper.GetString("smtp"),
		SmtpPort: viper.GetString("smtp_port"),
	})
	return &Bet{
		repo:       repo,
		mailSender: maiSender,
	}
}

/*
email: "test.adverts123@gmail.com"
password: "pUzWd4rUd7k8F5e"
smtp: "smtp.gmail.com"
smtp_port: ":587"
*/
