package main

import (
	"fmt"
	"mini-project-apotek/config"
)

func main() {
	config.InitDB()
	fmt.Print("halo")
}
