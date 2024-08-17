package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ZodiaSign struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name"`
	StartDate time.Time          `json:"start_date"`
	EndDate   time.Time          `json:"end_date"`
}
