package db

import (
	"context"
	"log"
	"my-echo/config"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dmManager struct {
	Ctx context.Context
	Db  *mongo.Database
}

var singletonDmManager *dmManager

var onceDmManger sync.Once

func GetDmManager() *dmManager {
	onceDmManger.Do(func() {
		log.Println("[INFO] Starting Initializing Singleton DB Manager")
		singletonDmManager = &dmManager{}
		singletonDmManager.initConnection()
	})
	return singletonDmManager
}

func (dm *dmManager) initConnection() {
	ctx := context.Background()
	dm.Ctx = ctx
	clientOpts := options.Client().ApplyURI(config.DatabaseConnectionString)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Println("[Error] Database initialize failed")
		log.Println("[Error]", err.Error())
		return
	}

	db := client.Database(config.DatabaseName)
	dm.Db = db

	log.Println("[INFO] Initialized Singleton DB Manager")
}
