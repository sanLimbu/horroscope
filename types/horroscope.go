package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Horoscope struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ZodiacSignID int                `json:"zodiac_sign_id"`
	Date         time.Time          `json:"date"`
	Type         string             `json:"type"`
	Prediction   string             `json:"prediction"`
}
