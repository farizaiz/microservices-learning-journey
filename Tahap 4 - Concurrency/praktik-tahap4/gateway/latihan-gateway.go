package main

import (
	"fmt"
	"sync"
	"time"
)

// MICROSERVICE 1: Ambil Profil (Butuh waktu lambat: 2 detik)
func getProfil(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	ch <- "👤 Profil: PT. Maju Mundur (Akun: Aktif)"
}

// MICROSERVICE 2: Ambil Tagihan (Butuh waktu sangat cepat: 0.5 detik)
func getTagihan(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	ch <- "💰 Tagihan: Rp 15.000.000 (Status: Belum Dibayar)"
}

// MICROSERVICE 3: Ambil Status Izin (Butuh waktu sedang: 1 detik)
func getStatusIzin(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	ch <- "📄 Izin: 2 Dokumen sedang dalam tahap Evaluasi Teknis"
}

func main() {
	fmt.Println("=== 🌐 Simulasi API Gateway Speed KLHK ===")
	waktuMulai := time.Now()

	var mesinAbsen sync.WaitGroup
	jalurData := make(chan string)

	// TODO 1: Daftarkan 3 tugas ke mesin absen
	// ... (Tulis kode Anda di sini) ...
	mesinAbsen.Add(3)

	// TODO 2: Panggil 3 fungsi microservice di atas menggunakan 'go'
	// Jangan lupa kirim alamat memory (&mesinAbsen) dan pipanya (jalurData)
	// ... (Tulis kode Anda di sini) ...

	go getProfil(&mesinAbsen, jalurData)
	go getTagihan(&mesinAbsen, jalurData)
	go getStatusIzin(&mesinAbsen, jalurData)

	// TODO 3: GOROUTINE MANDOR (TEKNIK BARU)
	// Kita membuat Goroutine tanpa nama (Anonymous Function) yang langsung dieksekusi.
	go func() {
		// 1. Mandor disuruh nungguin semua karyawan selesai
		mesinAbsen.Wait()

		// 2. Kalau selesai, Mandor menutup pipanya agar Bos berhenti menadah
		close(jalurData)
	}()

	fmt.Println("⏳ Menarik data dari 3 server berbeda secara serentak...")

	// TODO 4: TANGKAP DATA DENGAN FOR...RANGE
	// Buat perulangan 'for dataMasuk := range jalurData'.
	// Di dalamnya, cukup cetak 'dataMasuk' menggunakan fmt.Println.
	// Perulangan ini akan otomatis berhenti sendiri berkat perintah 'close(jalurData)' di atas!
	// ... (Tulis kode Anda di sini) ...

	for dataMasuk := range jalurData {
		fmt.Println(dataMasuk)
	}

	fmt.Printf("\n⚡ Semua data Dashboard berhasil dimuat dalam waktu: %v\n", time.Since(waktuMulai))
}
