package main

import (
	"fmt"
	"reflect"
)

/* Skenario: KLHK memiliki target pengurangan emisi karbon harian.
Anda diberikan sebuah Slice berisi data pengurangan emisi dari beberapa pabrik, dan sebuah angka target.
Temukan indeks dari dua pabrik yang jika dijumlahkan, hasilnya tepat sama dengan angka target tersebut! */

// ==========================================
// PENCARIAN PASANGAN TARGET EMISI
// ==========================================
func cariPasanganEmisi(emisiPabrik []int, target int) []int {
	// TODO: Tulis algoritma Anda di sini!
	// Hint: Gunakan perulangan bersarang (nested loop 'for' di dalam 'for'),
	// atau gunakan Map untuk solusi yang lebih cepat (O(n)).
	// Kembalikan slice berisi dua indeks pabrik yang memenuhi syarat
	sudahDilihat := make(map[int]int)

	for i, nilai := range emisiPabrik {
		pasangan := target - nilai

		if j, ada := sudahDilihat[pasangan]; ada {
			return []int{j, i}
		}
		sudahDilihat[nilai] = i
	}
	return []int{}
}

// ==========================================
// AREA TEST CASE (JANGAN UBAH BAGIAN INI)
// ==========================================
func main() {
	fmt.Println("=== 🚀 PENGUJIAN ALGORITMA SPEED KLHK ===")

	// --- TEST CASE ---
	fmt.Println("\n[Test 1] Cari Pasangan Emisi:")
	dataEmisi := []int{12, 5, 9, 2, 8}
	targetEmisi := 14

	hasil1 := cariPasanganEmisi(dataEmisi, targetEmisi)
	jawabanBenar1 := []int{1, 2} // Indeks 1 (nilai 5) + Indeks 2 (nilai 9) = 14

	// Validasi Test 1
	if reflect.DeepEqual(hasil1, jawabanBenar1) || reflect.DeepEqual(hasil1, []int{2, 1}) {
		fmt.Println("✅ LULUS! Algoritma pencarian Anda tepat sasaran.")
	} else {
		fmt.Printf("❌ GAGAL! Output Anda: %v | Seharusnya: %v\n", hasil1, jawabanBenar1)
	}

}
