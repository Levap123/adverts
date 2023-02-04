package handler

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	ErrInvalidTitle = fmt.Errorf("title min length - %d, title max length - %d",
		viper.GetInt("advert_title_min"), viper.GetInt("advert_title_max"))

	ErrInvalidBody = fmt.Errorf("body min length - %d, body max length - %d",
		viper.GetInt("advert_body_min"), viper.GetInt("body_max"))
)
