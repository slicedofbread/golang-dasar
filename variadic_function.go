package main

import "fmt"

func main() {
	// calling a variadic function with multiple arguments
	sum := addNumbers(1, 2, 3, 4, 5)
	fmt.Printf("The sum is: %d\n", sum)

	// calling a variadic function with a slice
	numbers := []int{10, 20, 30}
	sum2 := addNumbers(numbers...)
	fmt.Printf("The sum is: %d\n", sum2)
}

// variadic function that accepts any number of integer arguments
func addNumbers(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
