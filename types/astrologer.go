package types

import (
	"time"
)

// Astrologer represents the structure of the Astrologers table
type Astrologer struct {
	AstrologerID      int       `json:"astrologer_id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	PhoneNumber       string    `json:"phone_number"`
	Specialization    string    `json:"specialization"`
	ExperienceYears   int       `json:"experience_years"`
	Rating            float64   `json:"rating"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	Bio               string    `json:"bio"`
	AvailableFrom     time.Time `json:"available_from"`
	AvailableTo       time.Time `json:"available_to"`
}
