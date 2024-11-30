package main

import (
	boostrap "github.com/RaihanMalay21/api-service-riors/bootstrap"
	"github.com/RaihanMalay21/api-service-riors/config"
)

func main() {
	config.ConnectionDB()

	e := boostrap.SetupDependencies()

	e.Logger.Fatal(e.Start(":8080"))
}
