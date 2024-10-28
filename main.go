package main

import (
	"fmt"
	"ip-finder/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	application := app.Generate()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}