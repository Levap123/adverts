package validator

import "fmt"

func (v *Validator) IsTitleValid(title string) bool {
	fmt.Println(v)
	return len(title) >= v.advertTitleMin && len(title) <= v.advertTitleMax
}

func (v *Validator) IsBodyValid(body string) bool {
	return len(body) >= v.advertBodyMin && len(body) <= v.advertBodyMax
}

func (v *Validator) IsPriceValid(price int) bool {
	return price <= v.priceMax
}
