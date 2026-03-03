package main

import (
	"fmt"
)

func main() {
	// 1. Siapkan "tas ransel" kosong untuk menyimpan nama barang
	var daftarBarang []string
	var pilihan int
	var namaBarang string

	// Loop tanpa henti (Infinite Loop) agar aplikasi tidak langsung mati setelah pilih menu
	for {
		fmt.Println("\n=== 📦 Sistem Manajemen Inventaris ===")
		fmt.Println("1. Tambah Barang Baru")
		fmt.Println("2. Lihat Daftar Barang")
		fmt.Println("3. Keluar Aplikasi")
		fmt.Print("Masukkan pilihan menu (1/2/3): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Masukkan nama barang yang ingin ditambahkan: ")
			// Gunakan Scanln untuk mengambil teks dari user dan simpan ke variabel namaBarang
			fmt.Scanln(&namaBarang)

			// TODO (TUGAS ANDA):
			// Masukkan 'namaBarang' ke dalam slice 'daftarBarang' menggunakan perintah append!
			// ... (Tulis kode Anda di sini) ...

			daftarBarang = append(daftarBarang, namaBarang)

			fmt.Println("✅ Berhasil menambahkan", namaBarang, "ke dalam gudang!")

		case 2:
			fmt.Println("\n=== 📋 Daftar Barang di Gudang ===")

			// TODO (TUGAS ANDA):
			// Buat logika if-else: Jika len(daftarBarang) adalah 0, cetak "Gudang masih kosong".
			// Jika tidak kosong, gunakan for...range untuk mencetak semua barang beserta urutan angkanya.
			// ... (Tulis kode Anda di sini) ...

			if len(daftarBarang) == 0 {
				fmt.Println("Gudang masih kosong")
			} else {
				for index, value := range daftarBarang {
					fmt.Printf("%d %s\n", index+1, value)
				}
			}

		case 3:
			fmt.Println("Sampai jumpa! Menutup aplikasi...")
			return // Keluar dari program dan menghentikan loop

		default:
			fmt.Println("❌ Pilihan tidak valid. Silakan pilih 1, 2, atau 3.")
		}
	}
}
