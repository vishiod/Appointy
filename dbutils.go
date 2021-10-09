package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const CONNECTED = "Successfully connected to database: %v"

type MongoDatastore struct {
	db      *mongo.Database
	Session *mongo.Client
}

var mongoStore MongoDatastore = MongoDatastore{nil, nil}

func getDBStore() MongoDatastore {

	if mongoStore.db == nil {
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
	mongoDBStore.db = appDB
	mongoDBStore.Session = appDB.Client()
	return mongoDBStore
}
