package models

import "go.mongodb.org/mongo-driver/bson/primitive"




// Reader class
// Holds inherits the general user class
type Reader struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Fullname string	`json:"fullname" bson:"fullname"`
	Username string `json:"username" bson:"username"`
	Email string	`json:"email" bson:"email"`
	Password string 	`json:"password" bson:"password"`
	Books []Book `json:"books" bson:"books"`
	Active	bool	`json:"active" bson:"active"`
}

//ReaderEvent describes the event stored in the ICDE database
type ReaderEvent struct {
	BookTitle string `json:"book_title" bson:"book_title"`
	AuthorName string `json:"author_name" bson:"author_name"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Event events	`json:"event" bson:"event"`
}
