package validator

func (v *Validator) IsBetOk(bet int) bool {
	return bet >= 0
}
