package main

import "fmt"

func main() {
	var bulan = [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	slice1 := bulan[0:3]
	slice2 := bulan[3:6]

	fmt.Println(slice1)
	fmt.Println(slice2)

	//function len() untuk menghitung panjang slice
	fmt.Println("Panjang slice1:", len(slice1))
	fmt.Println("Panjang slice2:", len(slice2))

	//function cap() untuk menghitung kapasitas slice
	fmt.Println("Kapasitas slice1:", cap(slice1))
	fmt.Println("Kapasitas slice2:", cap(slice2))

	//FUNCTION APPEND() untuk menambahkan elemen ke dalam slice
	slice1 = append(slice1, "April")
	fmt.Println("Slice1 setelah ditambahkan elemen:", slice1)

	//function make() untuk membuat slice baru dengan panjang dan kapasitas tertentu
	slice3 := make([]string, 2, 5)
	slice3[0] = "Hello"
	slice3[1] = "World"
	fmt.Println("Slice3:", slice3)

	//function copy() untuk menyalin elemen dari satu slice ke slice lainnya
	slice4 := make([]string, len(slice1))
	copy(slice4, slice1)
	fmt.Println("Slice4 setelah disalin dari slice1:", slice4)

}
