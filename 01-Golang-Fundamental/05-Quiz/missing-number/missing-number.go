package main

import (
	"fmt"
)

/* Skenario: Staf KLHK sedang menyusun tumpukan berkas fisik yang sudah dinomori urut dari 0 sampai N.
Karena keteledoran, ada satu dokumen yang terselip dan hilang. Anda diberikan Slice berisi nomor-nomor dokumen yang tersisa.
Temukan nomor dokumen yang hilang tersebut! */

// ==========================================
// SOAL 3: MENCARI DOKUMEN YANG HILANG
// ==========================================
func cariDokumenHilang(antrean []int) int {
	n := len(antrean)
	totalSeharusnya := n * (n + 1) / 2
	totalAktual := 0

	for _, angka := range antrean {
		totalAktual += angka
	}

	return totalSeharusnya - totalAktual
}

// ==========================================
// AREA TEST CASE (JANGAN UBAH BAGIAN INI)
// ==========================================
func main() {
	fmt.Println("=== 🕵️ MESIN UJI LOGIKA SPEED KLHK ===")

	// --- TEST CASE SOAL 3 ---
	fmt.Println("\n[Test 3] Mencari Dokumen yang Hilang:")
	dataDokumen1 := []int{3, 0, 1}                   // Seharusnya ada 0, 1, 2, 3. Yang hilang: 2
	dataDokumen2 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1} // Yang hilang: 8

	hasil3a := cariDokumenHilang(dataDokumen1)
	hasil3b := cariDokumenHilang(dataDokumen2)

	if hasil3a == 2 && hasil3b == 8 {
		fmt.Println("✅ LULUS! Anda berhasil menemukan dokumen yang terselip dengan akurat.")
	} else {
		fmt.Printf("❌ GAGAL! Output Anda: (Test 1: %d, Test 2: %d) | Seharusnya: (Test 1: 2, Test 2: 8)\n", hasil3a, hasil3b)
	}
}
