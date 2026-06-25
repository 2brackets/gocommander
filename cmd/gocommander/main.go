package main

import (
	"log"

	"github.com/2brackets/gocommander/internal/ui"
)

func main() {
	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
