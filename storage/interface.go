package storage

import "6311_Project/models"

//Data
type InterfaceDB interface {
	AuthorSave(user interface{}, Collection string) (models.Author, error)
	FindByID(id interface{}, projection map[string]interface{}, result interface{}) error
	FindOne(collection string, key, pair string) (*models.Author, error)
	UpdateOneUser(fields map[string]interface{}, changes interface{}) error
	FindOneUser(collection string, key, pair string, result interface{}) error

	// Readers function
	ReaderSave(user interface{}, Collection string) (models.Reader, error)
	FindManyHistory(fields, projection, sort map[string]interface{}, limit, skip int64, results interface{}) error
	//book functions
	BookSave(user interface{}, Collection string) (models.Book, error)


	// event
	EventSave(user interface{}, Collection string, EventType bool) (interface{}, error)
	FindMany(fields, projection, sort map[string]interface{}, limit, skip int64, results interface{}) error
	FindManyEvents(fields, projection, sort map[string]interface{}, limit, skip int64,  results *[]map[string]interface{}) (error,[]map[string]interface{})

}
