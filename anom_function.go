package main

import "fmt"

type blacklist func(string) bool

func registerUser(name string, blacklist blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked")
	} else {
		fmt.Println("you are registered")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("bread", blacklist)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
