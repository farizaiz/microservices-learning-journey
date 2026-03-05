package main

import (
	"fmt"
)

// TODO 1: LENGKAPI FUNCTION INI
// Fungsi ini bertugas menghitung total tagihan.
// Parameter 1: 'pesanan' (slice of string) -> daftar makanan yang dipesan pelanggan
// Parameter 2: 'bukuMenu' (map[string]int) -> referensi harga makanan
// Return: Mengembalikan 2 nilai yaitu (totalItem, totalHarga)
func hitungTagihan(pesanan []string, bukuMenu map[string]int) (int, int) {
	totalItem := len(pesanan)
	totalHarga := 0

	// Gunakan perulangan (for _, makanan := range pesanan) untuk membaca satu per satu pesanan.
	// Lalu, ambil harga makanan tersebut dari 'bukuMenu' dan tambahkan ke 'totalHarga'.

	// ... (Tulis kode Anda di sini) ...

	for _, makanan := range pesanan {
		totalHarga += bukuMenu[makanan]
	}

	return totalItem, totalHarga // <-- Ubah agar mengembalikan totalItem dan totalHarga yang benar
}

func main() {
	// TODO 2: BUAT MAP (Buku Menu)
	// Buat map bernama 'menuRestoran' (Key: string, Value: int).
	// Isi langsung dengan minimal 3 menu, contoh: "Nasi Goreng": 20000, "Es Teh": 5000, dll.

	// ... (Tulis kode Anda di sini) ...
	menuRestoran := make(map[string]int)
	menuRestoran["NasiGoreng"] = 20000
	menuRestoran["MieKuah"] = 10000
	menuRestoran["EsTeh"] = 5000

	// TODO 3: BUAT SLICE (Kertas Pesanan)
	// Buat slice kosong bertipe string bernama 'pesananPelanggan'.
	// Slice ini akan digunakan untuk menyimpan makanan apa saja yang diketik oleh user.

	// ... (Tulis kode Anda di sini) ...

	var pesananPelanggan []string

	var pilihan int
	var namaPesanan string

	for {
		fmt.Println("\n=== 🍽️ Aplikasi Kasir Restoran ===")
		fmt.Println("1. Lihat Menu")
		fmt.Println("2. Tambah Pesanan")
		fmt.Println("3. Cetak Struk dan Bayar")
		fmt.Print("Pilih menu (1/2/3): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("\n=== 📖 Daftar Menu ===")
			// Gunakan for...range untuk mencetak semua isi 'menuRestoran'.
			// Contoh output: "Nasi Goreng - Rp 20000"
			for menu, harga := range menuRestoran {
				fmt.Printf("- %s : Rp %d\n", menu, harga)
			}

		case 2:
			fmt.Print("Masukkan nama makanan/minuman yang dipesan: ")
			fmt.Scanln(&namaPesanan)

			// TODO 4: VALIDASI DAN APPEND
			// a. Cek apakah 'namaPesanan' ada di dalam 'menuRestoran' menggunakan trik validasi Map.
			//    Contoh: _, tersedia := menuRestoran[namaPesanan]
			// b. Gunakan if-else:
			//    - Jika 'tersedia' adalah true, masukkan 'namaPesanan' ke dalam slice 'pesananPelanggan' menggunakan append(). Cetak pesan sukses.
			//    - Jika false, cetak pesan "Maaf, menu tersebut tidak ada di daftar".

			// ... (Tulis kode Anda di sini) ...

			_, tersedia := menuRestoran[namaPesanan]
			if tersedia {
				pesananPelanggan = append(pesananPelanggan, namaPesanan)
				fmt.Println("Pesanan tersedia di menu")
			} else {
				fmt.Println("Maaf, menu tersebut tidak ada di daftar")
			}

		case 3:
			fmt.Println("\n=== 🧾 STRUK PEMBAYARAN ===")

			// TODO 5: PANGGIL FUNCTION
			// a. Panggil fungsi 'hitungTagihan'.
			// b. Masukkan 'pesananPelanggan' dan 'menuRestoran' ke dalam tanda kurungnya.
			// c. Tangkap 2 nilai kembaliannya ke dalam variabel (misal: jumlah, total).

			// ... (Tulis kode Anda di sini) ...
			jumlah, total := hitungTagihan(pesananPelanggan, menuRestoran)

			// Hapus tanda komentar (//) di bawah ini jika TODO 5 sudah selesai:
			fmt.Printf("Total Item Pesanan : %d\n", jumlah)
			fmt.Printf("Total Pembayaran   : Rp %d\n", total)
			fmt.Println("Terima kasih atas kunjungan Anda!")
			return

		default:
			fmt.Println("❌ Pilihan tidak valid.")
		}
	}
}
