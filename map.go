package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":    "bread",
		"address": "123 Main Street",
		"email":   "bread@example.com",
	}

	fmt.Println("Name:", person["name"])
	fmt.Println("Address:", person["address"])
	fmt.Println("Email:", person["email"])
	fmt.Print(person)

	//len(person) //menghitung jumlah data pada map
	fmt.Println("Length:", len(person))

	//map mengambil data di map dengan key
	fmt.Println("Name:", person["name"])

	//map mengubah data di map dengan key
	person["name"] = "cake"
	fmt.Println("Updated Name:", person["name"])

	//make digunakan untuk membuat map baru
	newPerson := make(map[string]string)
	newPerson["name"] = "cookie"
	newPerson["address"] = "456 Elm Street"
	newPerson["email"] = "cookie@example.com"
	fmt.Println("New Person:", newPerson)

	//delete digunakan untuk menghapus data di map dengan key
	delete(person, "email")
	fmt.Println("After Deletion:", person)
}
