package main

import (
	m "go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/22_Middleware/Praktikum/Middlewares"
	"go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/22_Middleware/Praktikum/config"
	"go_Auliani-Sukma_Auliani-SUkma-Auliani-Sukma/22_Middleware/Praktikum/routes"

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