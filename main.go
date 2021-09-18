package main

import (
	"log"

	"github.com/nimerfarahty/go-rest/config"
	"github.com/nimerfarahty/go-rest/database"
	_ "github.com/nimerfarahty/go-rest/docs"
)

func main() {

	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	if err := database.Migrate(); err != nil {
		panic(err)
	}

	log.Fatal(runApp())

}
