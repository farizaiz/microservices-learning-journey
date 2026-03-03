package main

import "fmt"

func main() {
	var keranjang []string

	fmt.Println("Keranjang Awal:", keranjang)

	keranjang = append(keranjang, "Buku Golang")
	keranjang = append(keranjang, "Laptop", "Mouse")

	fmt.Println("Isi Keranjang Sekarang:", keranjang)

	totalBarang := len(keranjang)
	fmt.Println("Total barang bawaan ada:", totalBarang, "item")

	fmt.Println("\n=== Daftar Barang Bawaan ===")
	for urutan, barang := range keranjang {
		fmt.Printf("%d. %s\n", urutan+1, barang)
	}
}
