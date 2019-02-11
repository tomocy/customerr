package main

import (
	"log"
	"os"

	"github.com/tomocy/customerr/app"
)

func main() {
	app := app.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
