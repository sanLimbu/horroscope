package types

import "time"

type Horoscope struct {
	ZodiacSignID int       `json:"zodiac_sign_id"`
	Date         time.Time `json:"date"`
	Type         string    `json:"type"`
	Prediction   string    `json:"prediction"`
}
