package main

import "fmt"

type Dokumen struct {
	Perusahaan       string
	StatusLingkungan string
	Skor             int
}

// Fungsi algoritma Anda
func filterLulus(antrean []Dokumen) []string {
	// TODO 1: Buat variabel slice of string kosong untuk menampung nama perusahaan
	var hasil []string
	// TODO 2: Lakukan perulangan for...range pada 'antrean'
	for _, dok := range antrean {
		// TODO 3: Gunakan if untuk mengecek dua kondisi kelulusan menggunakan operator AND (&&)
		if dok.StatusLingkungan == "Aman" && dok.Skor >= 75 {
			// TODO 4: Jika lulus, append nama perusahaannya ke slice yang baru dibuat
			hasil = append(hasil, dok.Perusahaan)
		}
	}
	// TODO 5: Kembalikan (return) slice tersebut
	return hasil
}

func main() {
	antrean := []Dokumen{
		{Perusahaan: "PT Inovasi", StatusLingkungan: "Aman", Skor: 77},
		{Perusahaan: "PT Mandiri", StatusLingkungan: "Tidak Aman", Skor: 89},
		{Perusahaan: "PT Karya", StatusLingkungan: "Aman", Skor: 10},
		{Perusahaan: "PT Nusantara", StatusLingkungan: "Aman", Skor: 90},
	}

	lulus := filterLulus(antrean)
	fmt.Println("Perusahaan yang Lulus: ")
	for i, nama := range lulus {
		fmt.Printf("%d. %s\n", i+1, nama)
	}
}
