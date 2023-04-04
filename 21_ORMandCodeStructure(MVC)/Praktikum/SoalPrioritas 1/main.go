package main

import (
	"New_Praktikum/config"
	"New_Praktikum/router"
	
)

func main() {

	// Start Connection Database
	config.InitDB()

	// Routers
	e := router.Init()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
