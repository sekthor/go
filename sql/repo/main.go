package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	var conf Config

	file, err := os.ReadFile("./conf.yaml")
	if err != nil {
		log.Fatalf("could not open config file: %s", err.Error())
	}

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalf("could not read config file: %s", err.Error())
	}

	repo, err := NewRepository(conf)
	if err != nil {
		log.Fatalf("could not create repository: %s", err.Error())
	}

	err = repo.Connect()
	if err != nil {
		log.Fatalf("could not connect to database: %s", err.Error())
	}

	if err = repo.Ping(); err != nil {
		log.Fatalf("could not ping db: %s", err.Error())
	}
}
