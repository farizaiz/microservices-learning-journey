package main

import "fmt"

func main() {
	stokBarang := make(map[string]int)

	stokBarang["Laptop"] = 10
	stokBarang["Mouse"] = 25
	stokBarang["Buku Golang"] = 50

	fmt.Println("=== Informasi Stok Gudang ===")

	fmt.Println("Stok Khusus Mouse Saat Ini:", stokBarang["Mouse"], "unit")

	fmt.Println("\n=== Semua Daftar Stok ===")

	for nama, jumlah := range stokBarang {
		fmt.Printf("Barang %s | Sisa Stok: %d\n", nama, jumlah)
	}
}
