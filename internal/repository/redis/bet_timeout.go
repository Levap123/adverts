package goredis

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type BetTimeout struct {
	redis *redis.Client
}

func NewBetTimeout(redis *redis.Client) *BetTimeout {
	return &BetTimeout{
		redis: redis,
	}
}

func (bt *BetTimeout) ListenProducer(ch chan int, errCh chan error) {
	for i := 1; ; i++ {
		val, err := bt.redis.Get(context.TODO(), strconv.Itoa(i)).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				i = 0
				time.Sleep(time.Minute)
			} else {
				errCh <- err
				close(errCh)
				close(ch)
				return
			}
		}
		time, _ := strconv.Atoi(val)
		ch <- time
	}
}
