package types

type User struct {
	ID        string `bson:"_id" json,omitempty:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
