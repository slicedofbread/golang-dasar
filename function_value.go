package main

import "fmt"

func main() {
	goodBye := func(name string) string {
		return "Goodbye, " + name
	}

	fmt.Println(goodBye("bread"))
}
