package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. KITA SIAPKAN CETAKAN (STRUCT)
type BarangBukti struct {
	ID     string
	Nama   string
	Status string
}

// 2. KITA SIAPKAN KEMAMPUANNYA (POINTER METHOD)
func (b *BarangBukti) PeriksaLab() {
	// Simulasi waktu tunggu ke server Puslabfor (2 detik)
	time.Sleep(2 * time.Second)

	// Mengubah status asli di memori
	b.Status = "Selesai Uji Lab"
}

// 3. KITA BUAT PEKERJA GOROUTINE-NYA
// Fungsi ini menerima Pointer BarangBukti, Pointer WaitGroup, dan Pipa Channel
func jalankanPemeriksaan(bb *BarangBukti, wg *sync.WaitGroup, laporan chan string) {
	// TODO 1: Lapor ke "Mesin Absen" (WaitGroup) bahwa tugas ini SELESAI saat fungsi berakhir.
	// ... (Tulis kode Anda di sini) ...

	defer wg.Done()

	// TODO 2: Panggil method PeriksaLab() dari objek 'bb'
	// ... (Tulis kode Anda di sini) ...

	bb.PeriksaLab()

	// TODO 3: Rangkai pesan laporan keberhasilan menggunakan fmt.Sprintf
	// ... (Tulis kode Anda di sini) ...

	pesan := fmt.Sprintf("Barang %s (%s) berstatus: %s", bb.ID, bb.Nama, bb.Status)

	// TODO 4: Masukkan pesan tersebut ke dalam pipa 'laporan'
	// ... (Tulis kode Anda di sini) ...
	laporan <- pesan
}

func main() {
	fmt.Println("=== 🔬 Sistem Uji Lab Forensik Serentak ===")
	WaktuMulai := time.Now()

	// Simulasi tumpukan barang bukti dari database
	antreanBB := []BarangBukti{
		{ID: "BB-01", Nama: "Laptop Hitam", Status: "Menunggu Uji"},
		{ID: "BB-02", Nama: "Flashdisk 64GB", Status: "Menunggu Uji"},
		{ID: "BB-03", Nama: "Harddisk 1TB", Status: "Menunggu Uji"},
		{ID: "BB-04", Nama: "Handphone X", Status: "Menunggu Uji"},
	}

	var mesinAbsen sync.WaitGroup
	pipaLaporan := make(chan string)

	// 4. BOS MENYURUH KARYAWAN BEKERJA SECARA PARALEL
	for i := range antreanBB {

		// TODO 5: Tambahkan 1 tugas ke 'mesinAbsen'
		// ... (Tulis kode Anda di sini) ...

		mesinAbsen.Add(1)

		// Eksekusi pekerja sebagai Goroutine!
		// [TIPS EMAS]: Kita kirim alamat aslinya (&antreanBB[i]) agar Pointer Method-nya bekerja dengan benar.
		go jalankanPemeriksaan(&antreanBB[i], &mesinAbsen, pipaLaporan)
	}

	fmt.Println("⏳ Mengirim 4 barang bukti ke Puslabfor secara serentak...")

	// 5. BOS MENADAH PIPA LAPORAN
	// TODO 6: Buat perulangan 'for' yang berputar sebanyak jumlah data (len(antreanBB)).
	// Di dalamnya, tangkap pesan dari 'pipaLaporan' dan cetak menggunakan fmt.Println.
	// ... (Tulis kode Anda di sini) ...

	for i := 0; i < (len(antreanBB)); i++ {
		var laporanMasuk = <-pipaLaporan
		fmt.Printf("Bos menerima pesan %s\n", laporanMasuk)

	}

	// 6. BOS MENUNGGU SEMUA KARYAWAN ABSEN PULANG
	// TODO 7: Panggil method untuk menunggu semua WaitGroup selesai
	// ... (Tulis kode Anda di sini) ...

	mesinAbsen.Wait()

	fmt.Printf("\n🚀 Semua tugas selesai dalam waktu: %v\n", time.Since(WaktuMulai))
}
