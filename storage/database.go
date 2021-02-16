package storage

import (
	"6311_Project/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoconn struct {
	mongoConn 	*mongo.Client
	mongoDb		*mongo.Database
}

var secret, err = config.LoadSecrets()

func (m *mongoconn) mongoInit()  {
	connString := fmt.Sprintf("mongodb+srv://%v:%v@%v/<%v>?retryWrites=true&w=majority", secret.MONGO_USER, secret.MONGO_PASS, secret.MONGO_HOST, secret.MONGO_DB)
	client := options.Client()
	client.ApplyURI(connString)



}
