package models

import "go.mongodb.org/mongo-driver/bson/primitive"




// Reader class
// Holds inherits the general user class
type Reader struct {

}

//ReaderEvent describes the event stored in the ICDE database
type ReaderEvent struct {
	BookTitle string `json:"book_title" bson:"book_title"`
	AuthorName string `json:"author_name" bson:"author_name"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Event events	`json:"event" bson:"event"`
}
