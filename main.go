package main

import (
	"commandline/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("STARTED")

	aplication := app.Generate()
	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
