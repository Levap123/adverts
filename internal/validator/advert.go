package validator

func (v *Validator) IsAdvertValid(title, body string, price int) bool {
	return v.isTitleValid(title) && v.isBodyValid(body) && v.isPriceValid(price)
}

func (v *Validator) isTitleValid(title string) bool {
	return len(title) >= v.advertTitleMin && len(title) <= v.advertTitleMax
}

func (v *Validator) isBodyValid(body string) bool {
	return len(body) >= v.advertBodyMin && len(body) <= v.advertBodyMax
}

func (v *Validator) isPriceValid(price int) bool {
	return price <= v.priceMax
}
