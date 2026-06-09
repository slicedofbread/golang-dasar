package main

import (
	"fmt"
	"golang-dasar/database"
	_ "golang-dasar/internal" //underscore untuk menjalankan init internal tanpa harus memanggilnya secara langsung
)

func main() {
	fmt.Println(database.GetDatabase())
}
