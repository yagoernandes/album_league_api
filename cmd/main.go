package main

import (
	"log"

	"github.com/yagoernandes/album-league/cmd/environment"
)

func main() {

	server := environment.NewServer()

	SetRoutes(server)

	log.Fatal(server.App.Listen(":3000"))
}
