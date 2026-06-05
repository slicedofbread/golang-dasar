package main

import "fmt"

func goodBye(name string) string {
	return "good bye " + name
}

func main() {
	contoh := goodBye

	fmt.Println(contoh("bread"))

}
