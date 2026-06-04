package main

import "fmt"

func main() {
	var buah [3]string
	buah[0] = "Mangga"
	buah[1] = "Apel"
	buah[2] = "Jeruk"

	fmt.Println(buah[0]) // Output: Mangga
	fmt.Println(buah[1]) // Output: Apel
	fmt.Println(buah[2]) // Output: Jeruk

	// Deklarasi dan inisialisasi array secara langsung
	var angka = [5]int{1, 2, 3, 4, 5}
	fmt.Println(angka) // Output: [1 2 3 4 5]

	//length of array
	fmt.Println(len(buah))

	//mendapatkan data posisi index
	fmt.Println(buah[1]) // Output: Apel

	//mengubah data di posisi index
	buah[1] = "Pisang"
	fmt.Println(buah[1]) // Output: Pisang

}
