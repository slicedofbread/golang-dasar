package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b
	fmt.Println("Hasil penjumlahan: ", c)
	var d = a * b
	fmt.Println("Hasil perkalian: ", d)
	var e = b - a
	fmt.Println("Hasil pengurangan: ", e)
	var f = b / a
	fmt.Println("Hasil pembagian: ", f)

	//augmented assignment
	a += 5
	fmt.Println("Hasil augmented addition: ", a)
	b *= 2
	fmt.Println("Hasil augmented multiplication: ", b)

	//unary operator
	var z = 1
	z++
	fmt.Println("Hasil unary increment: ", z)
	z--
	fmt.Println("Hasil unary decrement: ", z)

	//oprasi perbandingan
	var isEqual = a == b
	fmt.Println("Apakah a sama dengan b? ", isEqual)
	var isGreater = a > b
	fmt.Println("Apakah a lebih besar dari b? ", isGreater)
	var isLess = a < b
	fmt.Println("Apakah a lebih kecil dari b? ", isLess)
	var isNotEqual = a != b
	fmt.Println("Apakah a tidak sama dengan b? ", isNotEqual)
	var isGreaterOrEqual = a >= b
	fmt.Println("Apakah a lebih besar atau sama dengan b? ", isGreaterOrEqual)
	var isLessOrEqual = a <= b
	fmt.Println("Apakah a lebih kecil atau sama dengan b? ", isLessOrEqual)

	//boolean operator
	var isTrue = true
	var isFalse = false
	var andResult = isTrue && isFalse
	fmt.Println("Hasil AND: ", andResult)
	var orResult = isTrue || isFalse
	fmt.Println("Hasil OR: ", orResult)
	var notResult = !isTrue
	fmt.Println("Hasil NOT: ", notResult)
}
