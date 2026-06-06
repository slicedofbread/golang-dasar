package main

import "fmt"

func main() {
	runApplication(true)
	fmt.Println("bread")
}

//defernya sambil baca panic pake recover
func endApplication() {
	fmt.Println("Aplikasi berhenti")
	message := recover()
	fmt.Println("Error message: ", message)
}

func runApplication(appError bool) {
	defer endApplication()
	if appError {
		panic("Aplikasi Error")
	}
}
