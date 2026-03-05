package main

import "fmt"

// 1. Membuat Cetakan Data (Struct)
// Biasakan nama Struct diawali dengan huruf kapital agar bisa diakses oleh file/microservice lainnya

type Barang struct {
	Nama     string
	Harga    int
	SisaStok int
}

func main() {
	// 2. Menggunakan Cetakan
	// Membuat variabel 'produk1' yang bertipe 'Barang' (bukan lagi string atau int)
	produk1 := Barang{
		Nama:     "Macbook",
		Harga:    1000000,
		SisaStok: 3,
	}
	produk2 := Barang{
		Nama:     "Meja",
		Harga:    1000000,
		SisaStok: 3,
	}

	fmt.Println("\n=== Info Produk 1 ===")
	fmt.Println("Nama Produk: ", produk1.Nama)
	fmt.Println("Harga Produk: Rp", produk1.Harga)
	fmt.Println("Stok: ", produk1.SisaStok)

	fmt.Println("\n=== Info Produk 2 ===")
	fmt.Printf("%s harganya Rp %d (Sisa: %d unit)\n", produk2.Nama, produk2.Harga, produk2.SisaStok)
}
