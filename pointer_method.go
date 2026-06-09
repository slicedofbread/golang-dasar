package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr " + man.Name
}

func main() {
	person := Man{"bread"}
	person.Married()

	fmt.Println(person.Name) // masih bread, karena method Married tidak merubah value, melainkan copy value

}
