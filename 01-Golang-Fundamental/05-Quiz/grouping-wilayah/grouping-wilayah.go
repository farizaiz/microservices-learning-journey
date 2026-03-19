package main

import (
	"fmt"
)

// ==========================================
// PENGELOMPOKAN WILAYAH (GROUPING)
// ==========================================
// Skenario: Anda menerima ratusan tumpukan dokumen evaluasi dari berbagai wilayah.
// Tugas Anda adalah mengelompokkan ID Dokumen tersebut ke dalam laci (Map) berdasarkan wilayahnya.

type Dokumen struct {
	ID               string
	Perusahaan       string
	Wilayah          string
	StatusLingkungan string
	Skor             int
}

func kelompokkanWilayah(daftarDokumen []Dokumen) map[string][]string {
	// TODO: Tulis algoritma Anda di sini!
	// Hint: Buat map dengan Key berupa string (Wilayah), dan Value berupa slice of string (kumpulan ID).
	// Lakukan perulangan pada daftarDokumen, lalu masukkan ID ke dalam map sesuai wilayahnya.

	hasilGroup := make(map[string][]string)
	for _, dok := range daftarDokumen {
		hasilGroup[dok.Wilayah] = append(hasilGroup[dok.Wilayah], dok.ID)
	}

	return hasilGroup
}

func main() {
	// --- TEST CASE ---
	fmt.Println("\n[Test 2] Kelompokkan Dokumen per Wilayah:")
	dataDokumen := []Dokumen{
		{ID: "DOC-1", Wilayah: "Jawa Barat"},
		{ID: "DOC-2", Wilayah: "Jawa Timur"},
		{ID: "DOC-3", Wilayah: "Jawa Barat"},
		{ID: "DOC-4", Wilayah: "Banten"},
		{ID: "DOC-5", Wilayah: "Jawa Timur"},
	}

	hasil2 := kelompokkanWilayah(dataDokumen)

	// Validasi Test 2 secara manual
	lulus := true
	if len(hasil2["Jawa Barat"]) != 2 || hasil2["Jawa Barat"][0] != "DOC-1" || hasil2["Jawa Barat"][1] != "DOC-3" {
		lulus = false
	}
	if len(hasil2["Jawa Timur"]) != 2 || hasil2["Jawa Timur"][0] != "DOC-2" || hasil2["Jawa Timur"][1] != "DOC-5" {
		lulus = false
	}
	if len(hasil2["Banten"]) != 1 || hasil2["Banten"][0] != "DOC-4" {
		lulus = false
	}

	if lulus {
		fmt.Println("✅ LULUS! Dokumen berhasil dikelompokkan dengan sempurna.")
	} else {
		fmt.Printf("❌ GAGAL! Pemetaan Anda keliru. Output Anda: %v\n", hasil2)
	}
}
