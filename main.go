package main

import (
	"comand-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Ponto de partida \n")

	application := app.Gerar()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
