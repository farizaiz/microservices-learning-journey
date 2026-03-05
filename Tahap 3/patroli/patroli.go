package main

import "fmt"

// TODO 1: BUAT STRUCT
// Buat struct bernama 'Pohon' dengan 3 properti (semua bertipe string):
// - Jenis  (misal: "Jati", "Mahoni")
// - Lokasi (misal: "Sektor-A", "Sektor-B")
// - Status (misal: "Sehat", "Sakit")
// ... (Tulis struct Anda di sini) ...

type Pohon struct {
	Jenis  string
	Lokasi string
	Status string
}

// TODO 2: BUAT METHOD DENGAN POINTER
// Buat method 'UbahStatus' yang menempel pada pointer *Pohon.
// Fungsi ini menerima 1 parameter 'statusBaru' (tipe string),
// lalu mengubah properti Status pada struct tersebut menjadi status yang baru.
// ... (Tulis method Anda di sini) ...

func (p *Pohon) UbahStatus(statusBaru string) {
	p.Status = statusBaru
}

// TODO 3: LENGKAPI FUNCTION DENGAN MAP
// Fungsi ini menerima daftar seluruh pohon, lalu mengembalikan rekap jumlah pohon di tiap lokasi.
func hitungStatistikArea(daftarPohon []Pohon) map[string]int {
	// Buat map kosong dengan Key string (lokasi) dan Value int (jumlah)
	statistik := make(map[string]int)

	// Lakukan perulangan for...range pada 'daftarPohon'.
	// Gunakan lokasi dari masing-masing pohon sebagai Key di dalam map,
	// lalu tambahkan nilainya (++).
	// ... (Tulis logika perulangan Anda di sini) ...
	for _, pohon := range daftarPohon {
		statistik[pohon.Lokasi]++
	}

	return statistik
}

func main() {
	// TODO 4: SIAPKAN SLICE UTAMA
	// Buat variabel 'dataHutan' yang merupakan Slice kosong dari struct Pohon.
	// ... (Tulis kode Anda di sini) ...

	var dataHutan []Pohon

	var pilihan int

	for {
		fmt.Println("\n=== 🌲 Sistem Pendataan Patroli Hutan ===")
		fmt.Println("1. Catat Pohon Baru")
		fmt.Println("2. Lihat Semua Data & Statistik Area")
		fmt.Println("3. Ubah Status Pohon (Pohon Pertama)")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var inputJenis, inputLokasi, inputStatus string
			fmt.Print("Jenis Pohon (Tanpa Spasi): ")
			fmt.Scanln(&inputJenis)
			fmt.Print("Lokasi (Tanpa Spasi, misal Sektor-A): ")
			fmt.Scanln(&inputLokasi)
			fmt.Print("Status (Sehat/Sakit): ")
			fmt.Scanln(&inputStatus)

			// TODO 5: KEMAS DAN SIMPAN (APPEND)
			// a. Buat variabel baru dari struct Pohon, isi dengan 3 data input di atas.
			// b. Masukkan ke dalam slice 'dataHutan' menggunakan append.
			// ... (Tulis kode Anda di sini) ...

			pohonBaru := Pohon{
				Jenis:  inputJenis,
				Lokasi: inputLokasi,
				Status: inputStatus,
			}

			dataHutan = append(dataHutan, pohonBaru)

			fmt.Println("✅ Data berhasil dicatat!")

		case 2:
			fmt.Println("\n=== 📋 Daftar Temuan Patroli ===")
			if len(dataHutan) == 0 {
				fmt.Println("Belum ada data.")
			} else {
				// Cetak daftar pohon
				for i, p := range dataHutan {
					fmt.Printf("%d. %s di %s (Status: %s)\n", i+1, p.Jenis, p.Lokasi, p.Status)
				}

				fmt.Println("\n=== 📊 Statistik Area ===")

				// TODO 6: PANGGIL FUNCTION MULTIPLE DATA
				// Panggil fungsi 'hitungStatistikArea' dan masukkan 'dataHutan' ke dalamnya.
				// Tampung hasilnya ke dalam variabel bernama 'hasilStatistik'.
				// ... (Tulis kode Anda di sini) ...

				hasilStatistik := hitungStatistikArea(dataHutan)

				// Hapus tanda komentar (//) di bawah ini jika TODO 6 selesai:
				for lokasi, jumlah := range hasilStatistik {
					fmt.Printf("- %s: %d pohon\n", lokasi, jumlah)
				}
			}

		case 3:
			if len(dataHutan) == 0 {
				fmt.Println("❌ Belum ada data untuk diubah!")
			} else {
				var statusBaru string
				fmt.Printf("Masukkan status baru untuk pohon %s di %s: ", dataHutan[0].Jenis, dataHutan[0].Lokasi)
				fmt.Scanln(&statusBaru)

				// TODO 7: EKSEKUSI METHOD POINTER
				// Panggil method UbahStatus() pada dataHutan[0] dengan parameter 'statusBaru'
				// ... (Tulis kode Anda di sini) ...
				dataHutan[0].UbahStatus(statusBaru)
				fmt.Println("✅ Status berhasil diperbarui!")
			}

		case 4:
			fmt.Println("Keluar dari sistem patroli...")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
