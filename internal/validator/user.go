package validator

import (
	"net/mail"
	"unicode"
)

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
