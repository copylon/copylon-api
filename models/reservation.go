package models

import (
	"time"
)

type Reservation struct {
	Intangible
	BookingTime       time.Time             `json:"bookingTime"`
	Broker            Organization          `json:"broker"`
	PriceCurrency     Currency              `json:"priceCurrency"`
	ReservationID     int                   `json:"reservationID" type:"int not null auto_increment"`
	ReservationStatus ReservationStatusType `json:"reservationStatus"`
	TotalPrice        int                   `json:"totalPrice"`
	UnderName         PersonOrOrganization  `json:"underName"`
}
