package models

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Address                 string  `json:"address"`
	AggregateRating         float32 `json:"aggregateRating"`
	BranchCode              string  `json:"branchCode"`
	Latitude                float64 `json:"latitude" validate:"latitude"`
	Logo                    string  `json:"logo" validate:"url"`
	Longitude               float64 `json:"longitude" validate:"longitude"`
	MaximumAttendeeCapacity int64   `json:"maximumAttendeeCapacity"`
	PhoneNumber             string  `json:"phoneNumber"`
}
