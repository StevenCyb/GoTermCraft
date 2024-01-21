package main

import (
	"ChromaTerm/pkg/terminal"
	"os"
)

func main() {
	terminal, err := terminal.New(os.Stdin, os.Stdout)
	if err != nil {
		terminal.Println("Error creating terminal:", err)
		return
	}

	width, height, err := terminal.Size()
	if err != nil {
		terminal.Println("Error getting terminal size:", err)
		return
	}

	terminal.Printf("Terminal size: %d columns x %d rows\n", width, height)
	terminal.Printf("name: %s\n", "steven")
}
