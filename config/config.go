package config

import (
	"log"
	"my-echo/enums"
	"os"

	"github.com/joho/godotenv"
)

var ServerPort string
var RunMode string

// Database
var DBServer string
var DBPort string
var DBUsername string
var DBPassword string
var DatabaseName string
var Database string
var DatabaseConnectionString string

// Auth
var PrivateKey string
var PublicKey string
var RegularTokenLifetime string

func InitEnvironment() {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = string(enums.DEVELOPMENT)
	}
	if RunMode != string(enums.PRODUCTION) {
		err := godotenv.Load()
		if err != nil {
			log.Println("[ERROR]: ", err.Error())
			return
		}
	}
	log.Println("RUN MODE:", RunMode)

	ServerPort = os.Getenv("SERVER_PORT")

	DBServer = os.Getenv("MONGO_SERVER")
	DBPort = os.Getenv("MONGO_PORT")
	DBUsername = os.Getenv("MONGO_USERNAME")
	DBPassword = os.Getenv("MONGO_PASSWORD")
	DatabaseName = os.Getenv("DATABASE_NAME")
	Database = os.Getenv("DATABASE")
	if Database == string(enums.MONGO) {
		DatabaseConnectionString = "mongodb://" + DBUsername + ":" + DBPassword + "@" + DBServer + ":" + DBPort
	}

	PrivateKey = os.Getenv("PRIVATE_KEY")
	PublicKey = os.Getenv("PUBLIC_KEY")
	RegularTokenLifetime = os.Getenv("REGULAR_TOKEN_LIFETIME")

}
