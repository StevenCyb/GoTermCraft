package main

import (
	"ChromaTerm/pkg/terminal"
	"fmt"
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
	terminal.Print("Test\n")
	terminal.Println("ok")
	input := ""
	for {
		char := terminal.ReadByte()
		if char != nil {
			terminal.Print(string(*char))
			input += string(*char)

			if *char == 13 {
				break
			} else if *char == 127 {
				input = input[:len(input)-2]
			}
		}
	}

	fmt.Println("Input:", input)
}
