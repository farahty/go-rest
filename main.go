package main

import (
	"log"
	"os"

	"github.com/nimerfarahty/go-rest/config"
	_ "github.com/nimerfarahty/go-rest/docs"
)

func main() {

	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	args := os.Args[1:]

	switch {

	case len(args) == 0:
		log.Fatal(runApp())
	case args[0] == "generate":
		Generate()
	case args[0] == "run":
		log.Fatal(runApp())
	default:
		log.Fatal("command not found")
	}

}
