package main

import (
	"fmt"
	"time"
)

// Ini adalah simulasi fungsi yang butuh waktu lama (misal: mengunduh file, atau query ke database)
func prosesDokumen(id int) {
	fmt.Printf("Mulai memvalidasi dokumen %d...\n", id)

	// time.Sleep memaksa program berhenti sejenak untuk mensimulasikan proses yang berat
	time.Sleep(2 * time.Second)

	fmt.Printf("✅ Dokumen %d selesai tervalidasi!\n", id)
}

func main() {
	fmt.Println("=== 🐢 CARA LAMA (BERURUTAN) ===")
	WaktuMulai1 := time.Now()

	// Memanggil fungsi secara normal (akan memakan waktu 2 + 2 + 2 = 6 detik)
	prosesDokumen(1)
	prosesDokumen(2)
	prosesDokumen(3)

	fmt.Printf("Total waktu cara lama: %v\n\n", time.Since(WaktuMulai1))

	fmt.Println("=== 🚀 CARA GOLANG (GOROUTINE) ===")
	WaktuMulai2 := time.Now()

	// TODO 1: Tambahkan kata 'go' di depan ketiga fungsi ini
	go prosesDokumen(4)
	go prosesDokumen(5)
	go prosesDokumen(6)

	// TODO 2: Tahan program utama agar tidak langsung tutup
	// Hapus tanda komentar pada baris time.Sleep di bawah ini nanti
	time.Sleep(3 * time.Second)

	fmt.Printf("Total waktu cara Golang: %v\n", time.Since(WaktuMulai2))
}
