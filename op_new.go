package main

import "fmt"

type address struct {
	city, province, country string
}

func main() {

	var alamat1 = new(address)
	var alamat2 = alamat1

	alamat2.city = "jakarta"

	fmt.Println(alamat1) // sudah jakarta
	fmt.Println(alamat2) // sudah jakarta

}
