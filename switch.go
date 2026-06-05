package main

import "fmt"

func main() {
	var name string = "bread"

	// switch statement
	switch name {
	case "bread":
		fmt.Println("Hello, bread!")
	case "cake":
		fmt.Println("Hello, cake!")
	default:
		fmt.Println("Hello, stranger!")
	}

	// switch dengan short statement
	switch length := len(name); {
	case length > 5:
		fmt.Println("Name is too long.")
	case length > 0:
		fmt.Println("Name is acceptable.")
	default:
		fmt.Println("Name cannot be empty.")
	}

	// switch tanpa kondisi (switch true)
	switch {
	case name == "bread":
		fmt.Println("Hello, bread!")
	case name == "cake":
		fmt.Println("Hello, cake!")
	default:
		fmt.Println("Hello, stranger!")
	}
}
