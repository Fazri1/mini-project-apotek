package main

import (
	"mini-project-apotek/config"
	"mini-project-apotek/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}