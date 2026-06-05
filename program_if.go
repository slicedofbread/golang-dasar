package main

import "fmt"

func main() {
	var name string = "bread"
	var age int = 20

	// if statement
	if name == "bread" {
		fmt.Println("Hello, bread!")
	}

	// if-else statement
	if age > 18 {
		fmt.Println("You are an adult.")
	} else if age >= 0 {
		fmt.Println("You are a minor.")
	} else {
		fmt.Println("Invalid age.")
	}

	// if dengan short statement
	if length := len(name); length > 5 {
		fmt.Println("Name is too long.")
	} else {
		fmt.Println("Name is acceptable.")
	}
}
