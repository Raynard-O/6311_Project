package storage

import (
	"6311_Project/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)




type mongoconn struct {
	mongoConn 	*mongo.Client
	mongoDb		*mongo.Database
	mongoCol	*mongo.Collection
}



var secret, err = config.LoadSecrets(false)
// mongoInit
// initialises database connection
func MongoInit(database , collection string, ctx context.Context) (InterfaceDB, error) {

	connString := fmt.Sprintf("mongodb+srv://%v:%v@%v/<%v>?retryWrites=true&w=majority", secret.MONGO_USER, secret.MONGO_PASS, secret.MONGO_HOST, secret.MONGO_DB)
	fmt.Println(connString)
	clientOptions := options.Client()
	clientOptions.ApplyURI(connString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	db := client.Database(database)
	mongoConn := mongoconn{
		mongoCol: db.Collection(collection),
		mongoConn: client,
		mongoDb:   db,
	}

	return &mongoConn, nil
}
