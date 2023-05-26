package main

import (
	"Code-Competence/config"
	"Code-Competence/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
