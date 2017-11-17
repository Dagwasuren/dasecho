package main

import (
	"log"

	"github.com/dasecho/dasecho/actions"
	"github.com/gobuffalo/envy"
)

func main() {
	addr := envy.Get("ADDR", ":3000")

	actions.WebHost = envy.Get("HOST", "https://dasecho.net")

	app := actions.App()
	app.Addr = addr
	log.Fatal(app.Serve())
}
