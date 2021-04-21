package storage

import (
	"6311_Project/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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




func (d *mongoconn) FindManyEvents(fields, projection, sort map[string]interface{}, limit, skip int64, results *[]map[string]interface{}) (error,[]map[string]interface{}) {
	ops := options.Find()
	if limit > 0 {
		ops.Limit = &limit
	}
	if skip > 0 {
		ops.Skip = &skip
	}
	if projection != nil {
		ops.Projection = projection
	}
	if sort != nil {
		ops.Sort = sort
	}

	cursor, err :=  d.mongoDb.Collection("readerEventType").Find(nil, fields, ops)

	//cursor, err := d.mongoDb.Collection("books").Find(nil, fields, ops)
	if err != nil {
		return err, nil
	}

	var output []map[string]interface{}
	for cursor.Next(nil) {
		var item map[string]interface{}
		_ = cursor.Decode(&item)
		output = append(output, item)
		//fmt.Println(item)
	}

	results = &output
	//if b, e := json.Marshal(output); e == nil {
	//	_ = json.Unmarshal(b, &results)
	//
	//	//fmt.Println(results)
	//} else {
	//	return e
	//}
	return nil, *results
}