package main

import (
	"fmt"
	"time"
)

// 1. KITA TAMBAHKAN PARAMETER CHANNEL (chan string)
func validasiCepat(id int, pipa chan string) {
	fmt.Printf("👷 Karyawan: Mulai cek dokumen %d...\n", id)

	// Pura-puranya butuh waktu 2 detik untuk ngecek database
	time.Sleep(2 * time.Second)

	// 2. KARYAWAN MENGIRIM PESAN KE DALAM PIPA
	// Gunakan fmt.Sprintf untuk merangkai teks tanpa langsung mencetaknya ke layar
	pesan := fmt.Sprintf("Lapor Bos! Dokumen %d dinyatakan VALID dan AMAN.", id)

	// Masukkan pesan tersebut ke dalam pipa!
	pipa <- pesan
}

func main() {
	fmt.Println("=== 📡 KOMUNIKASI GOROUTINE DENGAN CHANNEL ===")

	// 3. BOS MEMBUAT PIPA KOMUNIKASI
	// make(chan string) artinya pipa ini hanya bisa diisi oleh teks
	pipaLaporan := make(chan string)

	// 4. BOS MENYURUH 3 KARYAWAN BEKERJA (Sambil membekali mereka dengan pipanya)
	go validasiCepat(1, pipaLaporan)
	go validasiCepat(2, pipaLaporan)
	go validasiCepat(3, pipaLaporan)

	fmt.Printf("👨‍💼 Bos: Saya tunggu laporan kalian di ujung pipa...\n")

	// 5. BOS MENADAH PIPA 3 KALI (Sesuai jumlah dokumen/pekerja)
	// Kita pakai perulangan agar Bos stand-by menadah 3 kali berurut-urut
	for i := 1; i <= 3; i++ {

		// TODO 1: TANGKAP PESAN DARI PIPA
		// Buat variabel 'laporanMasuk', dan isi dengan data yang keluar dari 'pipaLaporan'.
		// Hint: gunakan tanda panah '<-' di depan nama pipanya.
		// ... (Tulis kode Anda di sini) ...

		var laporanMasuk = <-pipaLaporan

		// Hapus tanda komentar di bawah ini jika TODO 1 sudah dikerjakan
		fmt.Printf("📥 Bos menerima pesan: %s\n", laporanMasuk)
	}

	fmt.Println("\n👨‍💼 Bos: Semua dokumen sudah beres. Waktunya ngopi!")
}
