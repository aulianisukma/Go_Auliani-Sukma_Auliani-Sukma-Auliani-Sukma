package main

import (
	"Code-Competence/config"
	"Code-Competence/routes"
)

func main() {
	// start the server, and log if it fails
	config.InitDB()
	config.InitialMigration()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}