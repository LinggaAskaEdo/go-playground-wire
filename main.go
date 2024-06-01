package main

import (
	"log"
)

func main() {
	appConfig := NewAppConfig()

	scheduler, err := InitializedScheduler(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialize scheduler: %s\n", err)
	}

	scheduler.Start()

	server, err := InitializedServer(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialize server: %s\n", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Server is running")
}
