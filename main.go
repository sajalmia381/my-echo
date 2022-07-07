package main

import (
	"my-echo/config"
	"my-echo/db"
	"my-echo/src"
)

func main() {
	server := config.New()

	db.GetDmManager()

	src.Routes(server)

	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
