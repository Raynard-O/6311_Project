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
	EventID primitive.ObjectID `json:"_id" bson:"_id"`
	BookTitle string `json:"book_title" bson:"book_title"`
	Authors []string `json:"authors" bson:"authors"`
	Event Events	`json:"event" bson:"event"`
	EventType string `json:"event_type" bson:"event_type"`
}
