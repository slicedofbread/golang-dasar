package main

import "fmt"

func main() {
	conter := 0

	for conter <= 5 {
		fmt.Println("Counter:", conter)
		conter++
	}
	fmt.Println("Final Counter:")

	// for dengan short statement
	for i := 0; i <= 5; i++ {
		fmt.Println("i:", i)
	}
	fmt.Println("Final i")

	//init statement pada for
	for j := 0; j <= 5; j++ {
		fmt.Println("j:", j)
	}
	fmt.Println("Final j")

	//post statement pada for
	for k := 0; k <= 5; {
		fmt.Println("k:", k)
		k++
	}
	fmt.Println("Final k")

	//for range untuk iterasi pada array, slice, map, string, channel
	names := []string{"bread", "cake", "cookie"}
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}
	fmt.Println("Final names")

	// for dengan kondisi yang selalu benar (infinite loop)
	// for {
	// 	fmt.Println("This will run forever!")
	// }
}
