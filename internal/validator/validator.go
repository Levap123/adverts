package validator

import (
	"net/mail"
	"unicode"
)

type Validator struct {
	passwordMin int
	passwordMax int
	advertMin   int
	advertMax   int
	priceMax    int
}

func NewValidator(passwordMin, passwordMax, advertMin, advertMax, priceMax int) *Validator {
	return &Validator{
		passwordMin: passwordMin,
		passwordMax: passwordMax,
		advertMin:   advertMin,
		advertMax:   advertMax,
		priceMax:    priceMax,
	}
}

func (v *Validator) IsPasswordValid(password string) bool {
	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case !unicode.IsLetter(char) && !unicode.IsDigit(char):
			hasSpecial = true
		}
	}

	return len(password) >= v.passwordMin && len(password) <= v.passwordMax && hasUpper && hasLower && hasDigit && hasSpecial
}

func (v *Validator) IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
