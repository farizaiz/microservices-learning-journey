package main

import (
	"fmt"
	"sync" // <-- Tambahkan package sync
	"time"
)

// Kita tambahkan parameter kedua berupa POINTER ke WaitGroup
func prosesDokumen(id int, wg *sync.WaitGroup) {
	// wg.Done() bertugas melapor ke Bos bahwa tugas ini SELESAI.
	// Kita taruh di paling bawah agar dieksekusi terakhir.

	fmt.Printf("Mulai memvalidasi dokumen %d...\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("✅ Dokumen %d selesai tervalidasi!\n", id)

	wg.Done() // Karyawan absen pulang! (Mengurangi antrean 1)
}

func main() {
	fmt.Println("=== 🚀 CARA GOLANG PROFESIONAL (WAITGROUP) ===")
	WaktuMulai := time.Now()

	// 1. Bos membawa "Mesin Absensi"
	var mesinAbsen sync.WaitGroup

	// 2. Bos mencatat ada 3 dokumen yang harus diproses
	mesinAbsen.Add(3)

	// 3. Bos merekrut 3 karyawan (Goroutine)
	// Perhatikan tanda '&' (Pointer) agar mereka absen di mesin yang sama
	go prosesDokumen(4, &mesinAbsen)
	go prosesDokumen(5, &mesinAbsen)
	go prosesDokumen(6, &mesinAbsen)

	// 4. Bos menunggu (blocking) di sini sampai absen kembali ke 0
	mesinAbsen.Wait()

	// Baris di bawah ini TIDAK AKAN dieksekusi sampai Wait() di atas selesai
	fmt.Printf("Total waktu eksekusi: %v\n", time.Since(WaktuMulai))
}
