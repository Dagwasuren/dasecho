package main

import (
	"log"

	"github.com/dasecho/dasecho/actions"
	"github.com/gobuffalo/envy"
)

func main() {
	addr := envy.Get("ADDR", ":3000")
	app := actions.App()
	app.Addr = addr
	log.Fatal(app.Serve())
}
