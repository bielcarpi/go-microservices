package main

import (
	"broker/internal/server"
	"log"
)

const webPort = "80"

func main() {
	srv := server.NewServer()

	log.Printf("Starting broker service on port %s\n", webPort)

	if err := srv.Run(":" + webPort); err != nil {
		log.Panic(err)
	}
}
