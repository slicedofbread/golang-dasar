package main

import "fmt"

type address struct {
	city, province, country string
}

func main() {

	// pass by value
	address1 := address{"jakarta", "DKI Jakarta", "Indonesia"}
	adress2 := address1 // copy value, bukan reference
	adress2.city = "bandung"

	fmt.Println(address1) // masih jakarta
	fmt.Println(adress2)  // sudah bandung

	// pass by reference
	address3 := address{"jakarta", "DKI Jakarta", "Indonesia"}
	adress4 := &address3 // copy reference, bukan value
	adress4.city = "bandung"

	fmt.Println(address3) // sudah bandung
	fmt.Println(adress4)  // sudah bandung

}
