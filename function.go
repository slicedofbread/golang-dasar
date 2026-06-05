package main

import "fmt"

func main() {
	sayHello()
	result := add(5, 3)
	fmt.Printf("The result of adding 5 and 3 is: %d\n", result)
	sayhelloTo("bread", "butter")
	firstName, lastName := getFullName()
	fmt.Printf("Full Name: %s %s\n", firstName, lastName)
}

func sayHello() {
	fmt.Println("Hello, World!")
}

func add(a int, b int) int {
	return a + b
}

func sayhelloTo(firstName string, lastName string) {
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

// function returning multiple values
func getFullName() (string, string) {
	firstName := "davina"
	lastName := "karambol"
	return firstName, lastName
}
