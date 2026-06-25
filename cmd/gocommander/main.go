package main

import (
	"fmt"
	"os"

	"github.com/2brackets/gocommander/internal/filemanager"
)

func main() {
	fm := filemanager.New(".")

	entries, err := fm.ListDirectory()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		if entry.IsDir {
			fmt.Println("[DIR] ", entry.Name)
		} else {
			fmt.Println("      ", entry.Name)
		}
	}
}
