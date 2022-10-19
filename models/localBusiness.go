package models

type LocalBusiness struct {
	Organization
	CurrenciesAccepted []Currency    `json:"currenciesAccepted"`
	PaymentAccepted    []PaymentType `json:"paymentAccepted"`
}

