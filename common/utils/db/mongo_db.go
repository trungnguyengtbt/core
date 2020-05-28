package db

import (
	"context"
	"fmt"
	. "github.com/trungnguyengtbt/core/common/cons"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetMongoDBClient(conf MongoConfiguration) *mongo.Database {
	mgUri := conf.MgUri
	user := conf.MgUserName
	password := conf.MgPassword
	dbName := conf.MgDB
	authSource := conf.MgAuthSource

	fmt.Println("Connect mongo uri: ", mgUri)
	clientOptions := options.Client().ApplyURI(mgUri)
	credential := options.Credential{}
	credential.AuthMechanism = "SCRAM-SHA-1"
	if user != "" {
		fmt.Println("Connect mongo user: ", user)
		credential.Username = user
	}

	if password != "" {
		fmt.Println("Connect mongo password: ", password)
		credential.Password = password
	}

	if authSource != "" {
		fmt.Println("Connect mongo authSource: ", authSource)
		credential.AuthSource = authSource
	}
	// Set the credential clientOptions
	println("credential: ", credential)
	clientOptions.SetAuth(credential)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Create mongo client fail. ", err.Error())
		return nil
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Connect mongo fail. ", err.Error())
		return nil
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	clientDatabase := client.Database(dbName)
	return clientDatabase
}
