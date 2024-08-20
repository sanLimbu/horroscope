package types

import "time"

type Booking struct {
	BookingID       int       `json:"booking_id"`
	AstrologerID    int       `json:"astrologer_id"`
	UserID          int       `json:"user_id"`
	BookingDate     time.Time `json:"booking_date"`
	BookingTime     time.Time `json:"booking_time"`
	DurationMinutes int       `json:"duration_minutes"`
	Status          string    `json:"status"`
	Notes           string    `json:"notes"`
}
