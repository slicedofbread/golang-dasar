package main

import "fmt"

type address struct {
	city, province, country string
}

func ChangeCityToBandung(address *address) {
	address.city = "bandung"
}

func main() {
	alamat := &address{}
	ChangeCityToBandung(alamat)
	fmt.Println(alamat)
}
