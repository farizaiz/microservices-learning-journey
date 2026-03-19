package main

import "fmt"

// Fungsi algoritma Anda
func cariSpam(logID []string) []string {
	// TODO 1: Buat map kosong (Key: string, Value: int) untuk buku catatan frekuensi
	frekuensi := make(map[string]int)
	// TODO 2: Buat slice of string kosong untuk menampung ID yang terdeteksi spam
	var spam []string

	// TODO 3: Lakukan perulangan pertama pada 'logID' untuk MENGHITUNG.
	// Setiap ID yang lewat, tambahkan nilainya di dalam map (contoh: map[id]++)

	for _, id := range logID {
		frekuensi[id]++
	}

	// TODO 4: Lakukan perulangan kedua pada MAP yang sudah terisi.
	// Jika ada ID yang jumlahnya (value) > 1, append ID (key) tersebut ke slice spam

	for id, jumlah := range frekuensi {
		if jumlah > 1 {
			spam = append(spam, id)
		}
	}

	// TODO 5: Kembalikan (return) slice spam tersebut

	return spam
}

func main() {
	logID := []string{"A01", "B02", "A01", "C03", "B02", "B02"}
	hasil := cariSpam(logID)
	fmt.Println("ID Spam:")
	for i, id := range hasil {
		fmt.Printf("%d. %s\n", i+1, id)
	}

}
