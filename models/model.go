package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type user struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Fullname string	`json:"fullname" bson:"fullname"`
	Username string `json:"username" bson:"username"`
	Email string	`json:"email" bson:"email"`
}

type events struct {
	EventTime time.Time `json:"event_time" bon:"event_time"`
	EventID		primitive.ObjectID `json:"event_id" bson:"event_id"`
}
