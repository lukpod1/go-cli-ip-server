package main

import (
	"cli-ip-server/app"
	"log"
	"os"
)

func main() {
	application := app.Gerar()

	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
