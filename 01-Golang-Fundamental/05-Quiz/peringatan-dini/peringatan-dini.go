package main

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

/* Skenario: SPEED KLHK menerima ribuan aliran data dari sensor IoT (Internet of Things) yang dipasang di cerobong asap pabrik di seluruh Indonesia.
Anda diberikan sebuah Slice berisi data sensor udara (PM2.5).
Jika indeks PM2.5 lebih dari 100, pabrik tersebut masuk kategori "Berbahaya".
Karena datanya sangat masif, Anda wajib menggunakan Goroutine untuk mengecek setiap sensor secara serentak,
lalu mengembalikan Slice of String yang berisi nama-nama pabrik yang melanggar batas emisi.*/

type SensorIoT struct {
	NamaPabrik string
	PM25       int
}

// Fungsi pekerja — dijalankan tiap goroutine
func cekSensor(sensor SensorIoT, pipa chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	if sensor.PM25 > 100 {
		pipa <- sensor.NamaPabrik
	}
}

// ==========================================
// SOAL 5: DETEKSI POLUSI SERENTAK (CONCURRENCY)
// ==========================================
func deteksiPolusi(daftarSensor []SensorIoT) []string {
	var wg sync.WaitGroup

	// Gunakan channel untuk mengumpulkan nama pabrik yang melanggar
	pipaBahaya := make(chan string, len(daftarSensor))

	// TODO 1: Lakukan perulangan untuk menyebar Goroutine (Fan-Out)
	// Ingat untuk wg.Add(1) dan mengirim data ke fungsi pekerja
	for _, sensor := range daftarSensor {
		wg.Add(1)
		go cekSensor(sensor, pipaBahaya, &wg)
	}

	// TODO 2: Buat Goroutine Mandor (Anonymous Function) untuk menunggu wg.Wait() lalu close(channel)
	go func() {
		wg.Wait()
		close(pipaBahaya)
	}()

	// TODO 3: Tangkap data dari channel menggunakan for...range (Fan-In)
	// Siapkan sebuah slice of string kosong, lalu append data dari channel ke slice tersebut
	var hasilBahaya []string
	for namaPabrik := range pipaBahaya {
		hasilBahaya = append(hasilBahaya, namaPabrik)
	}

	// TODO 4: Kembalikan (return) slice tersebut
	return hasilBahaya // Ganti dengan slice hasil Anda
}

// ==========================================
// AREA TEST CASE (JANGAN UBAH BAGIAN INI)
// ==========================================
func main() {
	fmt.Println("=== 🌪️ RADAR POLUSI SPEED KLHK ===")

	dataHariIni := []SensorIoT{
		{NamaPabrik: "PT. Langit Biru", PM25: 45},
		{NamaPabrik: "PT. Asap Hitam", PM25: 150}, // Berbahaya!
		{NamaPabrik: "PT. Hijau Daun", PM25: 80},
		{NamaPabrik: "PT. Limbah Udara", PM25: 210}, // Berbahaya!
		{NamaPabrik: "PT. Nafas Sesak", PM25: 105},  // Berbahaya!
		{NamaPabrik: "PT. Angin Segar", PM25: 12},
	}

	// Eksekusi fungsi buatan Anda
	hasil := deteksiPolusi(dataHariIni)

	// Kita urutkan sesuai abjad agar mudah divalidasi oleh sistem
	// (karena Goroutine sifatnya acak)
	sort.Strings(hasil)

	jawabanBenar := []string{"PT. Asap Hitam", "PT. Limbah Udara", "PT. Nafas Sesak"}
	sort.Strings(jawabanBenar)

	fmt.Printf("Pabrik Terdeteksi: %v\n", hasil)

	if reflect.DeepEqual(hasil, jawabanBenar) {
		fmt.Println("✅ LULUS!")
	} else {
		fmt.Printf("❌ GAGAL! Output Anda tidak sesuai.\n")
	}
}
