package thirdparty

import (
	"6311_Project/models"
	"6311_Project/storage"
	"context"
	"fmt"
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




func getAllAcceptedEvents( authorsName string) (error, []map[string]interface{}) {
	DB, err := storage.MongoInit("ICDE", "books", context.Background())

	if err != nil {
		return err, nil
	}

	if err != nil {
		return err, nil
	}
	field := make(map[string]interface{})
	filter := make(map[string]interface{})
	filter["$regex"] = authorsName

	field["authors"] = filter

	var output,ou  []map[string]interface{}
	fmt.Println(field)
	err, ou = DB.FindManyEvents(field, nil, nil, 0, 0, &output)
	//fmt.Println( ou)
	if err != nil {
		return err,nil
	}
	return nil, ou
}



func Recommend(c echo.Context, authorsName string) error {
	DB, err := storage.MongoInit("ICDE", "readers", context.Background())
	err, events := getAllAcceptedEvents(authorsName)
	if err != nil {
		return err
	}


	for _,v := range events{
		result := new(models.Reader)

		err = DB.FindByID(v["user_id"], nil, &result )

		fields := make(map[string]interface{})
		fields["_id"] = v["user_id"]

		changes := make(map[string]interface{})
		message := "New Message!!! \n  " + "New Book from: " + v["authors"].(string) + "\n" + "Book Title: " + v["book_title"].(string) + "\n"
		changes["notification"] = append(result.Notification, message)
		//send them a  document
		err = DB.UpdateOneUser(fields, changes)
	if err != nil {
		return err
	}

	}

	return c.JSONPretty(http.StatusOK, "lastEvent", "")
}
