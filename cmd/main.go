package main

import (
	"log"

	"github.com/yagoernandes/album_league_api/cmd/environment"
)

func main() {

	server := environment.NewServer()

	SetRoutes(server)

	log.Fatal(server.App.Listen(":3000"))
}
