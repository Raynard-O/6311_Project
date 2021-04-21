package storage

import (
	"6311_Project/models"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *mongoconn) BookSave(user interface{}, Collection string) (models.Book, error) {
	var output models.Book
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


func (d *mongoconn) FindMany(fields, projection, sort map[string]interface{}, limit, skip int64, results interface{}) error {
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
//oplogReplay: true,
//	filter: { ts: { $gte: new Timestamp(1514764800, 0) } } }

	cursor, err :=  d.mongoDb.Collection("books").Find(nil, fields, ops)

	//cursor, err := d.mongoDb.Collection("books").Find(nil, fields, ops)
	if err != nil {
		return err
	}

	var output []map[string]interface{}
	for cursor.Next(nil) {
		var item map[string]interface{}
		_ = cursor.Decode(&item)
		output = append(output, item)
		//fmt.Println(output)
	}

	if b, e := json.Marshal(output); e == nil {
		_ = json.Unmarshal(b, &results)

		//fmt.Println(results)
	} else {
		return e
	}
	return nil
}