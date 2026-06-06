package main

import "fmt"

func sayhellowithfilter(name string, filter func(string) string) {
	fmt.Println("hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "****"
	} else {
		return name
	}
}

func main() {
	sayhellowithfilter("bread", spamFilter)

	pake := spamFilter
	sayhellowithfilter("anjing", pake)
}
