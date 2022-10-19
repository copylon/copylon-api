package models

import (
	"time"
)

type Organization struct {
	Thing
	Address         string        `json:"address"`
	Awards          []string      `json:"awards"`
	Department      *Organization `json:"department"`
	DissolutionDate time.Time     `json:"dissolutionDate"`
	Email           string        `json:"email, omitempty" validate:"email, omitempty"`
	FaxNumber       string        `json:"faxNumber"`
	Founder         *Person       `json:"person"`
	FoundingDate    time.Time     `json:"foundingDate"`
	LegalName       string        `json:"legalName"`
	Location        *Location     `json:"location"`
	Member          *Organization `json:"member,omitempty"`
	TaxID           string        `json:"taxID"`
	Telephone       string        `json:"telephone"`
}

func (organization Organization) getName() string {
	return organization.Name
}

func (organization Organization) getEmail() string {
	return organization.Email
}
