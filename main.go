package main

import (
	"log"
	"os"

	"github.com/joshmedeski/sesh/seshcli"
)

func main() {
	app := seshcli.App()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
