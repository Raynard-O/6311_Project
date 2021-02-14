package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	BookID string `json:"book_id" bson:"book_id"`
	Content string	`json:"content" bson:"content"`
}

// Author class
// Holds inherits the general user class
type Author struct {
	User user `json:"user" bson:"user"`
	Books []Book `json:"books" bson:"books"`
}

//AuthorEvent describes the event stored in the ICDE database
type AuthorEvent struct {
	BookTitle string `json:"book_title" bson:"book_title"`
	AuthorID primitive.ObjectID `json:"author_id" bson:"author_id"`
	Event events	`json:"event" bson:"event"`
}
