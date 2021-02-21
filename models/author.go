package models

import "go.mongodb.org/mongo-driver/bson/primitive"



// Author class
// Holds inherits the general user class
type Author struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Fullname string	`json:"fullname" bson:"fullname"`
	Username string `json:"username" bson:"username"`
	Email string	`json:"email" bson:"email"`
	Password string 	`json:"password" bson:"password"`
	User User `json:"user" bson:"user"`
	Books []Book `json:"books" bson:"books"`
	Active	bool	`json:"active" bson:"active"`
}

//AuthorEvent describes the event stored in the ICDE database
type AuthorEvent struct {
	BookTitle string `json:"book_title" bson:"book_title"`
	AuthorID primitive.ObjectID `json:"author_id" bson:"author_id"`
	Event events	`json:"event" bson:"event"`
}
