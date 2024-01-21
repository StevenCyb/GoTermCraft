package main

import (
	"ChromaTerm/pkg/terminal/form"
	"fmt"
)

func main() {
	// username, _ := form.NewInput().
	// 	OnValidate(func(input string) bool {
	// 		return len(input) >= 3
	// 	}).
	// 	Label("Username").
	// 	Read(false)
	password, _ := form.NewInput().
		IsSecret().
		OnValidate(func(input string) bool {
			return len(input) >= 8
		}).
		Label("Password").
		Read(false)

	// fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
}
