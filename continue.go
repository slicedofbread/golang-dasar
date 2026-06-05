package main

import "fmt"

func main() {
	// continues breaks out of the innermost loop, while break can be used to exit from any loop.
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke =", i)
	}
	fmt.Println("Finished the loop")
}
