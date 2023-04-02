package main

import (
	"Praktikum/config"
	"Praktikum/config"
)

func main (){
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":2424"))

}