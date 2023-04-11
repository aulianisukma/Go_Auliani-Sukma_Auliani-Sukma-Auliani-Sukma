package main

import (
	"go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/23_UnitTesting/Praktikum/config"

	"go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/23_UnitTesting/Praktikum/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitDB()
	// create a new echo instance
	e := routes.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}