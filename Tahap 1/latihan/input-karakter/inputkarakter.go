package main

import "fmt"

func main() {
	var angka1, angka2 int

	fmt.Print("Masukkan Angka Pertama: ")
	fmt.Scanln(&angka1)

	fmt.Print("Masukkan Angka Kedua: ")
	fmt.Scanln(&angka2)

	hasil := angka1 + angka2
	fmt.Println("Hasil Penjumlahan adalah : ", hasil)
}
