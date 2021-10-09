package dbutils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDatastore struct {
	DB      *mongo.Database
	Session *mongo.Client
}

var mongoStore MongoDatastore = MongoDatastore{nil, nil}

func GetDBStore() MongoDatastore {

	if mongoStore.DB == nil {
		mongoStore = createDBStore()
	}
	return mongoStore
}

func createDBStore() MongoDatastore {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	appDB := client.Database("mydb")

	var mongoDBStore MongoDatastore
	mongoDBStore.DB = appDB
	mongoDBStore.Session = appDB.Client()
	return mongoDBStore
}
