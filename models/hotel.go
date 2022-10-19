package models

import (
	"time"
)

type Hotel struct {
	LocalBusiness
	CheckInTime   time.Time `json:"checkInTime"`
	CheckOutTime  time.Time `json:"checkOutTime"`
	NumberOfRooms int32     `json:"numberOfRooms"`
	PetsAllowed   bool      `json:"petsAllowed"`
	StarRating    *Rating   `json:"starRating"`
}
