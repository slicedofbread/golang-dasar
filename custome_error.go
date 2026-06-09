package main

import (
	"fmt"
)

// customError adalah tipe error kustom.
type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}

func checkAge(age int) error {
	if age < 18 {
		return &customError{message: "umur tidak boleh di bawah 18"}
	}
	return nil
}

func main() {
	err := checkAge(16)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Umur valid")
}
