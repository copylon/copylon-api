package models

type PaymentType struct {
	Thing
	Name          string `json:"name"`
	AlternateName string `json:"alternateName"`
}
