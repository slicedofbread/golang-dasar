package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil fungsi")
}

func run() {
	// defer akan tetap di eksekusi seltelah fungsi yang manggilnya selesai walaupun terjadi error
	defer logging()
	fmt.Println("menjalankan aplikasi")
}

func main() {
	run()
}
