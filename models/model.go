package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password []byte             `json:"password,omitempty" bson:"password,omitempty"`
	Role     string             `json:"role" bson:"role"`
}
type Attendance struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Date     string             `json:"Date" bson:"Date"`
	ClockIn  string             `json:"ClockIn" bson:"ClockIn"`
	ClockOut string             `json:"ClockOut" bson:"ClockOut"`
	Role     string             `json:"role" bson:"role"`
}
