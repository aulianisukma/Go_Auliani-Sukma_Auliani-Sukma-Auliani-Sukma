package main

import (
	m "Praktikum/Middlewares"
	"Praktikum/config"
	"Praktikum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitDB()
	// create a new echo instance
	e := routes.New()

	m.LogMiddlewares(e)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}