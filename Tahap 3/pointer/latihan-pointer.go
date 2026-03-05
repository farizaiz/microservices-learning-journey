package main

import "fmt"

type Barang struct {
	Nama     string
	Harga    int
	SisaStok int
}

// INI ADALAH METHOD DENGAN POINTER
// Perhatikan (b *Barang) sebelum nama fungsi. Ini yang membuatnya "menempel" ke Struct Barang.
// Tanda bintang (*) artinya kita menggunakan Pointer untuk mengubah harga aslinya.
func (b *Barang) BeriDiskon(persen int) {
	fmt.Printf("\n💸 Memproses diskon %d%% untuk %s...\n", persen, b.Nama)

	// Rumus diskon: (Harga * Persen) / 100
	potongan := (b.Harga * persen) / 100

	// Kurangi harga asli dengan potongan
	b.Harga = b.Harga - potongan
}

func main() {
	produk1 := Barang{
		Nama:     "Laptop Mac",
		Harga:    15000000,
		SisaStok: 5,
	}

	fmt.Println("=== Info Produk Sebelum Diskon ===")
	fmt.Println("Nama  :", produk1.Nama)
	fmt.Println("Harga : Rp", produk1.Harga)

	// KITA MEMANGGIL METHOD-NYA DI SINI
	// Perhatikan betapa elegan cara memanggilnya, menggunakan tanda titik (.)
	produk1.BeriDiskon(10) // Memberikan diskon 10%

	fmt.Println("\n=== Info Produk Setelah Diskon ===")
	fmt.Println("Nama  :", produk1.Nama)
	fmt.Println("Harga : Rp", produk1.Harga) // Harganya akan berubah secara permanen!
}
