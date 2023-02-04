package validator

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
