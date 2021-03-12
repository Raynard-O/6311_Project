package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Book struct {
	BookID primitive.ObjectID `json:"_id" bson:"_id"`
	BookName string	`json:"book_name" bson::"book_name"`
	Pages int8	`json:"pages" bson:"pages"`
	Authors []string `json:"authors" bson:"authors"`
	UploadDate time.Time `json:"upload_date" bson:"upload_date"`
	Content string	`json:"content" bson:"content"`
}


type User struct {
	Fullname string	`json:"fullname" bson:"fullname"`
	Username string `json:"username" bson:"username"`
	Email string	`json:"email" bson:"email"`
	Password string 	`json:"password" bson:"password"`
}

type Events struct {
	EventTime time.Time `json:"event_time" bon:"event_time"`
}
