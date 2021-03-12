package storage

import (
	"6311_Project/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *mongoconn) EventSave(user interface{}, Collection string, EventType bool) (interface{}, error) {
	if EventType{
		var output models.AuthorEvent
		_result, err := d.mongoDb.Collection(Collection).InsertOne(context.Background(), user)

		if err != nil {
			return output, err
		}

		err = d.mongoDb.Collection(Collection).FindOne(nil, bson.M{"_id": _result.InsertedID}).Decode(&output)

		if err != nil {
			return output, err
		}
		return output, nil
	}else {
		var output models.ReaderEvent
		_result, err := d.mongoDb.Collection(Collection).InsertOne(context.Background(), user)

		if err != nil {
			return output, err
		}

		err = d.mongoDb.Collection(Collection).FindOne(nil, bson.M{"_id": _result.InsertedID}).Decode(&output)

		if err != nil {
			return output, err
		}
		return output, nil
	}

}

