package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Levap123/adverts/configs"
	"github.com/Levap123/adverts/internal/mail"
	"github.com/Levap123/adverts/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Bet struct {
	repo       repository.BetRepo
	repoAdvert repository.AdvertRepo
	mailSender *mail.MailSender
}

func NewBet(repo repository.BetRepo, repoAdvert repository.AdvertRepo) *Bet {
	conf := configs.EmailConf{
		Email:    viper.GetString("email"),
		Password: viper.GetString("password"),
		Smtp:     viper.GetString("smtp"),
		SmtpPort: viper.GetInt("smtp_port"),
	}
	maiSender := mail.NewMailSender(conf)
	return &Bet{
		repo:       repo,
		repoAdvert: repoAdvert,
		mailSender: maiSender,
	}
}

func (b *Bet) MakeBet(ctx context.Context, userId, advertId, betPrice int) (int, error) {
	email, err := b.repoAdvert.GetEmail(ctx, advertId)
	if err != nil {
		return 0, fmt.Errorf("service - make bet - %w", err)
	}

	ok, err := b.repo.IsActive(ctx, advertId)
	if err != nil {
		return 0, fmt.Errorf("service - make bet - %w", err)
	}

	if !ok {
		return 0, fmt.Errorf("service - make bet - %w", ErrAdvertIsNotActive)
	}

	advertPrice, err := b.repo.GetAdvertPrice(ctx, advertId)
	if err != nil {
		return 0, fmt.Errorf("service - make bet - %w", err)
	}
	if betPrice < advertPrice {
		return 0, fmt.Errorf("service - make bet - %w", ErrPriceSmall)
	}

	priceCurrent, err := b.repo.GetPrice(ctx, advertId)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("service - make bet - %w", err)
		}
		betId, err := b.repo.Create(ctx, userId, advertId, betPrice)
		if err != nil {
			return 0, fmt.Errorf("service - make bet - %w", err)
		}
		if err := b.mailSender.Send(email, fmt.Sprintf("your advert number %d is under bet with the price %d!",
			advertId, betPrice)); err != nil {
			logrus.Errorf("service - make bet - mail - %v", err)
		}
		return betId, nil
	}
	if betPrice <= priceCurrent {
		return 0, fmt.Errorf("service - make bet - %w", ErrPriceSmall)
	}

	betId, err := b.repo.Update(ctx, userId, advertId, betPrice)
	if err != nil {
		return 0, fmt.Errorf("service - make bet - %w", err)
	}

	if err := b.mailSender.Send(email, fmt.Sprintf("your advert number %d is under bet with the price %d!",
		advertId, betPrice)); err != nil {
		logrus.Errorf("service - make bet - mail - %v", err)
	}
	return betId, nil
}

/*
email: "test.adverts123@gmail.com"
password: "pUzWd4rUd7k8F5e"
smtp: "smtp.gmail.com"
smtp_port: ":587"
*/
