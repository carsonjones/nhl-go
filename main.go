package main

import (
	"go-nhl/cmd"
	"log"
)

func main() {
	config := cmd.NewConfig()
	config.ParseFlags()

	if err := config.Execute(); err != nil {
		log.Fatal(err)
	}
}
