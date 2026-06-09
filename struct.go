package main

import "fmt"

type Customer struct {
	Name    string
	address string
	Age     int
}

// struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, " my name is ", customer.Name)
}

func main() {

	var bread Customer
	bread.Name = "bread"
	bread.address = "jakarta"
	bread.Age = 20

	var jaka Customer
	jaka.Name = "jaka"
	jaka.address = "bandung"
	jaka.Age = 25

	fmt.Println(bread)

	// memanggil atribut struct
	fmt.Println(bread.Name)
	fmt.Println(bread.address)
	fmt.Println(bread.Age)

	// struct literal langsung
	robert := Customer{
		Name:    "robert",
		address: "bandung",
		Age:     25,
	}
	fmt.Println(robert)

	budi := Customer{"budi", "surabaya", 30}
	fmt.Println(budi)

	// memanggil method struct
	bread.sayHello("robert")
	jaka.sayHello("budi")
}
