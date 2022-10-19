package models

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	gorm.Model
	Address     string    `json:"address"`
	BirthDate   time.Time `json:"birthDate"`
	BirthPlace  Place     `json:"birthPlace"`
	Email       string    `json:"email" validate:"email"`
	FamilyName  string    `json:"familyName"`
	Gender      Gender    `json:"gender"`
	GivenName   string    `json:"givenName"`
	MiddleName  string    `json:"middleName"`
	Nationality string    `json:"nationality"`
	PhoneNumber string    `json:"phoneNumber"`
}

func (person Person) getName() string {
	return person.GivenName + person.FamilyName
}

func (person Person) getEmail() string {
	return person.Email
}
