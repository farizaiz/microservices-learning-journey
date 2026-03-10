package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. STRUCT (Cetakan Dokumen Evaluasi SPEED)
type LaporanEvaluasi struct {
	ID               string
	Perusahaan       string
	StatusLingkungan string // "Aman", "Waspada", "Kritis"
	Tervalidasi      bool   // Awalnya false
}

// 2. POINTER METHOD
// TODO 1: Buat method 'Validasi' untuk pointer *LaporanEvaluasi
// Method ini bertugas mengubah properti 'Tervalidasi' menjadi true.
// ... (Tulis method Anda di sini) ...

func (p *LaporanEvaluasi) Validasi() {
	p.Tervalidasi = true
}

// 3. FUNGSI PEKERJA (GOROUTINE)
// Menerima data laporan, menunda 1 detik (simulasi cek database), memvalidasi, dan mengirim ke pipa
func prosesEvaluasi(doc *LaporanEvaluasi, wg *sync.WaitGroup, ch chan LaporanEvaluasi) {
	defer wg.Done()

	// Simulasi sistem sedang mengecek indikator lingkungan ke server pusat
	time.Sleep(1 * time.Second)

	// TODO 2: Panggil method untuk memvalidasi dokumen ini
	// ... (Tulis kode Anda di sini) ...
	doc.Validasi()

	// TODO 3: Kirim dokumen yang SUDAH divalidasi tersebut (bukan pointernya, tapi wujud aslinya: *doc) ke dalam channel 'ch'
	// ... (Tulis kode Anda di sini) ...
	ch <- *doc
}

func main() {
	fmt.Println("=== 🍃 SPEED KLHK: Evaluasi Dokumen Serentak ===")
	waktuMulai := time.Now()

	// 4. SLICE (Database Antrean SPEED)
	daftarDokumen := []LaporanEvaluasi{
		{ID: "DOC-01", Perusahaan: "PT. Tambang Jaya", StatusLingkungan: "Waspada", Tervalidasi: false},
		{ID: "DOC-02", Perusahaan: "PT. Sawit Makmur", StatusLingkungan: "Aman", Tervalidasi: false},
		{ID: "DOC-03", Perusahaan: "PT. Kertas Nusantara", StatusLingkungan: "Waspada", Tervalidasi: false},
		{ID: "DOC-04", Perusahaan: "PT. Limbah Kimia", StatusLingkungan: "Kritis", Tervalidasi: false},
		{ID: "DOC-05", Perusahaan: "PT. Hutan Lestari", StatusLingkungan: "Aman", Tervalidasi: false},
	}

	var wg sync.WaitGroup

	// Pipa ini membawa tipe data Struct LaporanEvaluasi utuh!
	pipaHasil := make(chan LaporanEvaluasi, len(daftarDokumen))

	// 5. DISTRIBUSI TUGAS KE GOROUTINE
	fmt.Println("⏳ Memproses", len(daftarDokumen), "dokumen evaluasi secara serentak...")
	for i := range daftarDokumen {
		wg.Add(1)
		// Ingat: Kirim alamat memori aslinya (&) agar method pointer-nya bisa bekerja
		go prosesEvaluasi(&daftarDokumen[i], &wg, pipaHasil)
	}

	// 6. MANDOR PENGAMAT (API Gateway Pattern)
	go func() {
		wg.Wait()
		close(pipaHasil)
	}()

	// 7. MAP UNTUK STATISTIK
	// Kita ingin tahu ada berapa perusahaan yang status lingkungannya Aman/Waspada/Kritis
	statistik := make(map[string]int)

	// 8. MENANGKAP HASIL & MENGHITUNG STATISTIK (FOR...RANGE)
	for docMasuk := range pipaHasil {

		// TODO 4: Cetak ID, Nama Perusahaan, dan status Tervalidasi dari dokumen yang masuk
		// Contoh: "✅ DOC-01 (PT. Tambang Jaya) selesai dievaluasi: true"
		// ... (Tulis kode Anda di sini) ...
		fmt.Printf("%s (%s) selesai dievaluasi: %t\n", docMasuk.ID, docMasuk.Perusahaan, docMasuk.Tervalidasi)

		// TODO 5: Tambahkan hitungan ke dalam map 'statistik' berdasarkan StatusLingkungan-nya
		// ... (Tulis kode Anda di sini) ...
		statistik[docMasuk.StatusLingkungan]++
	}

	// 9. CETAK HASIL AKHIR (MAP)
	fmt.Println("\n=== 📊 Statistik Status Lingkungan KLHK ===")
	// TODO 6: Lakukan for...range pada map 'statistik' untuk mencetak hasilnya
	// ... (Tulis kode Anda di sini) ...
	for status, jumlah := range statistik {
		fmt.Printf("- Status %s: %d Perusahaan\n", status, jumlah)
	}

	fmt.Printf("\n⚡ Waktu Total: %v\n", time.Since(waktuMulai))
}
