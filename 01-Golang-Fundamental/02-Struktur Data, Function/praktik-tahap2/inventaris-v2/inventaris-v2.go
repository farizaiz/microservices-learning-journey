package main

import (
	"fmt"
)

// 1. INI ADALAH FUNCTION DENGAN MULTIPLE RETURN VALUES
// Fungsi ini menerima data 'inventaris' (map), lalu mengembalikan 2 angka (int, int): (totalBarang, totalAset)
func hitungStatistik(inventaris map[string]int) (int, int) {
	// TODO 1 (TUGAS ANDA):
	// a. Buat variabel untuk menghitung jumlah barang (gunakan len() pada map 'inventaris').
	// b. Buat variabel 'totalAset' dengan nilai awal 0.
	// c. Lakukan perulangan (for _, harga := range inventaris) untuk menjumlahkan semua 'harga' ke dalam 'totalAset'.
	// d. Kembalikan (return) kedua variabel tersebut secara bersamaan.

	// ... (Tulis kode Anda di sini) ...
	totalBarang := len(inventaris)
	totalAset := 0

	for _, harga := range inventaris {
		totalAset = totalAset + harga
	}

	return totalBarang, totalAset // <-- Ganti baris ini dengan return variabel yang benar hasil hitungan Anda
}

func main() {
	// TODO 2 (TUGAS ANDA):
	// Buat variabel bernama 'daftarBarang'.
	// Variabel ini harus berupa Map dengan Key bertipe 'string' (Nama) dan Value bertipe 'int' (Harga).
	// Petunjuk: Gunakan perintah make(...)

	// ... (Tulis kode Anda di sini) ...

	daftarBarang := make(map[string]int)

	var pilihan int
	var namaBarang string
	var hargaBarang int

	// Loop tanpa henti (Infinite Loop)
	for {
		fmt.Println("\n=== 📦 Sistem Manajemen Inventaris V2 (Map & Function) ===")
		fmt.Println("1. Tambah Barang Baru")
		fmt.Println("2. Lihat Daftar Barang & Statistik")
		fmt.Println("3. Keluar Aplikasi")
		fmt.Print("Masukkan pilihan menu (1/2/3): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Masukkan nama barang: ")
			fmt.Scanln(&namaBarang)
			fmt.Print("Masukkan harga barang: ")
			fmt.Scanln(&hargaBarang)

			// TODO 3 (TUGAS ANDA):
			// Masukkan 'hargaBarang' ke dalam map 'daftarBarang' menggunakan kunci 'namaBarang'.

			// ... (Tulis kode Anda di sini) ...

			daftarBarang[namaBarang] = hargaBarang

			fmt.Println("✅ Berhasil menambahkan", namaBarang, "seharga Rp", hargaBarang)

		case 2:
			fmt.Println("\n=== 📋 Daftar Barang di Gudang ===")

			// Kita anggap daftarBarang sudah berhasil Anda buat di TODO 2
			if len(daftarBarang) == 0 {
				fmt.Println("Gudang masih kosong")
			} else {
				// Map tidak punya index otomatis, jadi kita buat variabel urutan sendiri
				urutan := 1
				for nama, harga := range daftarBarang {
					fmt.Printf("%d. %s - Rp %d\n", urutan, nama, harga)
					urutan++
				}

				// TODO 4 (TUGAS ANDA):
				// Panggil fungsi 'hitungStatistik' dan masukkan 'daftarBarang' ke dalam tanda kurungnya.
				// Tangkap 2 nilai yang dikembalikan fungsi tersebut ke dalam variabel baru,
				// misalnya beri nama variabelnya: jumlahItem, totalNilai

				// ... (Tulis kode Anda di sini) ...
				jumlahItem, totalNilai := hitungStatistik(daftarBarang)

				// Jika TODO 4 sudah dikerjakan, HAPUS tanda komentar (//) pada dua baris di bawah ini:
				fmt.Println("-----------------------------------")
				fmt.Printf("📊 STATISTIK: Ada %d jenis barang dengan Total Aset Rp %d\n", jumlahItem, totalNilai)
			}

		case 3:
			fmt.Println("Sampai jumpa! Menutup aplikasi...")
			return

		default:
			fmt.Println("❌ Pilihan tidak valid.")
		}
	}
}
