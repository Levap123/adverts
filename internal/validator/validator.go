package validator

import (
	"fmt"

	"github.com/Levap123/adverts/configs"
)

type Validator struct {
	passwordMin    int
	passwordMax    int
	advertBodyMin  int
	advertBodyMax  int
	advertTitleMin int
	advertTitleMax int
	priceMax       int
}

func NewValidator(confs configs.ValidatorConf) *Validator {
	fmt.Println(confs)
	return &Validator{
		passwordMin:    confs.PasswordMin,
		passwordMax:    confs.PasswordMax,
		advertBodyMin:  confs.AdvertBodyMin,
		advertTitleMin: confs.AdvertTitleMin,
		advertTitleMax: confs.AdvertTitleMax,
		advertBodyMax:  confs.AdvertBodyMax,
		priceMax:       confs.PriceMax,
	}
}
