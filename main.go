package main

import (
	"ip-finder/app"
	"log"
	"os"
)

func main() {

	if err := app.Generate().Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}