package main

import "fmt"

type address struct {
	city, province, country string
}

func main() {
	address3 := address{"jakarta", "DKI Jakarta", "Indonesia"}
	adress4 := &address3 // copy reference, bukan value
	adress4.city = "bandung"

	fmt.Println(address3) // sudah bandung
	fmt.Println(adress4)  // sudah bandung

	// adress4 = &address{"surabaya", "Jawa Timur", "Indonesia"}
	*adress4 = address{"surabaya", "Jawa Timur", "Indonesia"}
	fmt.Println(address3) // masih bandung
	fmt.Println(adress4)  // sudah surabaya
}
