package main

import (
	"log"

	"github.com/teach310/genta"
)

func main() {
	if err := genta.Run(); err != nil {
		log.Fatalln(err)
	}
}
