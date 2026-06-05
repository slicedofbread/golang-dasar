package main

import "fmt"

func main() {

	// The break statement is used to exit a loop prematurely when a certain condition is met.
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("Exited the loop")

	// The continue statement is used to skip the current iteration of a loop and move on to the next iteration. It does not exit the loop entirely, but rather skips the remaining code in the current iteration and proceeds to the next one. In contrast, break exits the loop immediately, regardless of where it is in the loop. Therefore, while break can be used to exit from any loop, continue only affects the current iteration of the innermost loop it is placed in.
}
