package helper

var version = "1.0.0"
var Application = "Belajar Golang"

func sayGoodBye(name string) string {
	return "Good bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

// untuk bisa di akses huruf depan harus kapital, kalau tidak maka hanya bisa di akses di dalam package helper saja, tidak bisa di akses di luar package helper
// kalau huruf depan kecil maka hanya bisa di akses di dalam package helper saja, tidak bisa di akses di luar package helper
