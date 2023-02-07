package service

import (
	"fmt"
	"time"

	"github.com/Levap123/adverts/internal/repository"
)

type BetTimeout struct {
	repoBet    repository.BetRepo
	repoAdvert repository.AdvertRepo
	producer   repository.BetTimeoutRepo
}

func NewBetTimeout(repoBet repository.BetRepo, repoAdvert repository.AdvertRepo, producer repository.BetTimeoutRepo) *BetTimeout {
	return &BetTimeout{
		repoBet:    repoBet,
		repoAdvert: repoAdvert,
		producer:   producer,
	}
}

func (bt *BetTimeout) ListenConsumer(ch chan int, errCh chan error) error {
	go bt.producer.ListenProducer(ch, errCh)

	for {
		select {
		case err := <-errCh:
			return fmt.Errorf("servuce - listen consumer - %w", err)
		case dur := <-ch:
			if dur <= int(time.Now().Unix()) {
				/* TODO
				1. CHANGE ADVERT STATUS
				2. DELETE BET
				*/
			}
		}
	}
}
