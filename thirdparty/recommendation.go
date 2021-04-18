package thirdparty

import (
	"6311_Project/controllers"
	"6311_Project/storage"
	"context"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)


var err error

type thirdParty struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	lastId primitive.ObjectID `json:"last_id" bson:"last_id"`
	Name string	`json:"name" bson:"name"`
}




func (t *thirdParty)recommend(c echo.Context) error {

	DB, err := storage.MongoInit("ICDE", "authors", context.Background())
	
	if err != nil {
		return controllers.BadRequestResponse(c, err.Error())
	}
	
	last := thirdParty{
		ID:     primitive.NewObjectID(),
		lastId: primitive.ObjectID{},
		Name: "lastID",
	}

	lastEvent, err := DB.AuthorSave(last, "3rdParty")
	//get the last event you saw

	//get everyother event after
	//filter the author events
	//get reader who searched for the this author
	//send new author events to reader
	return c.JSONPretty(http.StatusOK, lastEvent, "")
}
