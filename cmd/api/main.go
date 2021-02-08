package main

import (
	"go-tednica/internal/platform/di"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	return di.ProvideServer().Run()
}
