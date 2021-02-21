package storage

import (
	"6311_Project/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * Save
 * Save is used to save a record in the mongoStore
 */

func (d *mongoconn) AuthorSave(user interface{}, Collection string) (models.Author, error) {
	var output models.Author
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

/**
 * Find One by
 */
func (d *mongoconn) FindOne(collection string, key, pair string) (*models.Author, error) {
	result := new(models.Author)
	filter := bson.M{key: pair}
	//ops := options.FindOne()
	//ops.Projection = projection
	if err := d.mongoDb.Collection(collection).FindOne(nil, filter).Decode(&result); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.New("document not found")
		}
		return nil, err
	}
	return result, nil
}


/**
 * FindByID
 * find a single record by id in the mongoStore
 * returns nil if record isn't found.
 *
 * param: interface{}            id
 * param: map[string]interface{} projection
 * return: map[string]interface{}
 */
func (d *mongoconn) FindByID(id interface{}, projection map[string]interface{}, result interface{}) error {
	ops := options.FindOne()
	if projection != nil {
		ops.Projection = projection
	}
	if err := d.mongoCol.FindOne(nil, bson.M{"_id": id}, ops).Decode(result); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("document not found")
		}
		return err
	}
	return nil
}

/**
 * UpdateOne
 *
 * Updates one item in the mongoStore using fields as the criteria.
 *
 * param: map[string]interface{} fields
 * param: interface{}            payload
 * return: error
 */
func (d *mongoconn) UpdateOneUser(fields map[string]interface{}, changes interface{}) error {
	var u map[string]interface{}
	if err := d.mongoCol.FindOneAndUpdate(nil, fields, bson.M{
		"$set": changes,
	}).Decode(&u); err != nil {
		return err
	}
	return nil
}

/**
* Find One by
 */
func (d *mongoconn) FindOneUser(collection string, key, pair string, result interface{}) error {

	filter := bson.M{key: pair}

	//ops.Projection = projection
	if err := d.mongoDb.Collection(collection).FindOne(nil, filter).Decode(result); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return errors.New("document not found")
		}
		return err
	}
	return nil
}
