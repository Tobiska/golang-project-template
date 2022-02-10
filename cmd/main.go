package main

import (
	"golang-project-template/config"
	"golang-project-template/internal/app"
	"log"
	//"golang-project-template/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
