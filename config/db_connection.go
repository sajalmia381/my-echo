package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client
var isConnectedToDb bool

func InitDBConnection() bool {

	// Create a new client and connect to the server
	log.Println("[INFO] Creating database connection")
	client, err := createConnection()
	if err != nil {
		isConnectedToDb = false
		return false
	}
	mongoClient = client

	// defer func() {
	// 	if err := mongoClient.Disconnect(context.TODO()); err != nil {
	// 		log.Println("[ERROR] while disconnect: ", err)
	// 	}
	// 	log.Println("[INFO] Database connection closed")
	// }()
	isConnectedToDb = true
	log.Println("[INFO] DB connection initialized")
	return true
}

func createConnection() (*mongo.Client, error) {
	_client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DatabaseConnectionString))
	return _client, err
}

func ping() bool {
	// Ping the primary
	if isConnectedToDb {
		if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
			log.Println("[ERROR] while pinging: .", err)
			return false
		}
		// log.Println("[INFO] Successfully connected and pinged.")
		return true
	} else {
		log.Println("[ERROR] DB Connection doesn't exist")
	}
	return false
}

func reconnect() {
	log.Println("[INFO] Trying to reconnect to DB")
	for !isConnectedToDb {
		time.Sleep(time.Duration(1) * time.Second)
		log.Println("[INFO] Trying to reconnect to DB")
	}
	isConnectedToDb = true
}

func DBHealthChecker() {
	log.Println("[INFO] Starting db health checker")
	for {
		if !ping() {
			isConnectedToDb = false
			break
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
	reconnect()
}
