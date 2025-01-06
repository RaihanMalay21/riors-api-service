package main

import (
	boostrap "github.com/RaihanMalay21/api-service-riors/bootstrap"
	"github.com/RaihanMalay21/api-service-riors/config"
)

func main() {
	config.ConnectionDB()
	config.ConnectionRedis()
	config.GoogleOuth20()

	e := boostrap.SetupDependencies()

	e.Logger.Fatal(e.Start(":8080"))
}
