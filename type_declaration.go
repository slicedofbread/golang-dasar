package main

func main() {

	// saya mengubah tipe data string menjadi NoKTP dengan cara membuat type baru dengan nama NoKTP
	type NoKTP string

	// saya membuat variable dengan tipe data NoKTP
	var ktpSaya NoKTP = "1234567890"
	println(ktpSaya)

	// saya juga bisa mengubah tipe data string menjadi NoKTP dengan cara casting
	var contohktp string = "0987654321"
	var contohktp2 NoKTP = NoKTP(contohktp)
	println(contohktp2)
}
