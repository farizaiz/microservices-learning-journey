package main

import "fmt"

// TODO 1: BUAT STRUCT 'Pengajuan'
// Properti:
// - Nomor      (string) -> misal: "IZIN-001"
// - Perusahaan (string) -> misal: "PT.Maju"
// - Jenis      (string) -> misal: "Tambang", "Hutan", "Air"
// - Status     (string) -> selalu berawal dari "Pending"
// ... (Tulis struct Anda di sini) ...

type Pengajuan struct {
	Nomor      string
	Perusahaan string
	Jenis      string
	Status     string
}

// TODO 2: BUAT METHOD POINTER
// Buat method bernama 'Setujui' yang menempel pada *Pengajuan.
// Fungsi ini tidak menerima parameter apa-apa,
// hanya bertugas mengubah properti Status milik pengajuan tersebut menjadi "Disetujui".
// ... (Tulis method Anda di sini) ...
func (p *Pengajuan) Setujui() {
	p.Status = "Disetujui"
}

// TODO 3: FUNCTION STATISTIK
// Buat fungsi 'hitungStatistikJenis' yang menerima slice dari Pengajuan,
// lalu mengembalikan map[string]int berisi jumlah pengajuan berdasarkan 'Jenis'-nya.
// ... (Tulis fungsi Anda di sini) ...

func hitungStatistikJenis(databaseIzin []Pengajuan) map[string]int {
	statistik := make(map[string]int)

	for _, pengajuan := range databaseIzin {
		statistik[pengajuan.Jenis]++
	}
	return statistik
}

func main() {
	// TODO 4: SLICE DATABASE
	// Buat slice kosong dari struct Pengajuan bernama 'databaseIzin'
	// ... (Tulis kode Anda di sini) ...
	var databaseIzin []Pengajuan

	var pilihan int

	for {
		fmt.Println("\n=== 📝 Sistem Manajemen Izin KLHK ===")
		fmt.Println("1. Buat Pengajuan Baru")
		fmt.Println("2. Lihat Semua Pengajuan & Statistik")
		fmt.Println("3. Setujui Pengajuan (Cari berdasarkan Nomor)")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var inNomor, inPT, inJenis string
			fmt.Print("Nomor Izin (Tanpa Spasi, misal IZIN-001): ")
			fmt.Scanln(&inNomor)
			fmt.Print("Nama Perusahaan (Tanpa Spasi): ")
			fmt.Scanln(&inPT)
			fmt.Print("Jenis Izin (Tambang/Hutan/Air): ")
			fmt.Scanln(&inJenis)

			// TODO 5: APPEND DATA BARU
			// Buat struct Pengajuan baru dengan data di atas.
			// Ingat: Set Status awal secara manual menjadi "Pending".
			// Lalu masukkan ke 'databaseIzin'.
			// ... (Tulis kode Anda di sini) ...

			pengajuanBaru := Pengajuan{
				Nomor:      inNomor,
				Perusahaan: inPT,
				Jenis:      inJenis,
				Status:     "Pending",
			}

			databaseIzin = append(databaseIzin, pengajuanBaru)

			fmt.Println("✅ Pengajuan berhasil didaftarkan dengan status Pending!")

		case 2:
			fmt.Println("\n=== 📋 Daftar Pengajuan ===")
			if len(databaseIzin) == 0 {
				fmt.Println("Belum ada data.")
			} else {
				// Menampilkan daftar
				for i, p := range databaseIzin {
					fmt.Printf("%d. [%s] %s - %s (Status: %s)\n", i+1, p.Nomor, p.Perusahaan, p.Jenis, p.Status)
				}

				fmt.Println("\n=== 📊 Statistik Jenis Izin ===")
				// TODO 6: TAMPILKAN STATISTIK
				// Panggil fungsi 'hitungStatistikJenis', tangkap hasilnya, lalu cetak menggunakan for...range.
				// ... (Tulis kode Anda di sini) ...

				hasilStatistik := hitungStatistikJenis(databaseIzin)

				for jenis, jumlah := range hasilStatistik {
					fmt.Printf("- %s: %d Pengajuan\n", jenis, jumlah)
				}
			}

		case 3:
			if len(databaseIzin) == 0 {
				fmt.Println("❌ Belum ada pengajuan untuk disetujui!")
			} else {
				var targetNomor string
				fmt.Print("Masukkan Nomor Izin yang ingin disetujui: ")
				fmt.Scanln(&targetNomor)

				ditemukan := false

				// TODO 7: LOGIKA PENCARIAN & UPDATE (LEVEL HARD)
				// Gunakan perulangan 'for' (bisa pakai range atau for biasa) untuk mengecek isi 'databaseIzin'.
				// Jika ada pengajuan yang Nomor-nya sama (==) dengan 'targetNomor':
				//   a. Panggil method Setujui() HANYA untuk pengajuan tersebut.
				//   b. Ubah variabel 'ditemukan' menjadi true.
				//   c. Gunakan 'break' agar perulangan langsung berhenti setelah ketemu.

				// [TIPS EMAS GOLANG]:
				// Jika menggunakan 'for index, data := range databaseIzin',
				// pastikan Anda memanggil method-nya dari slice aslinya: databaseIzin[index].Setujui()
				// JANGAN menggunakan data.Setujui() karena itu hanya akan mengubah data fotokopiannya!

				// ... (Tulis kode Anda di sini) ...
				for index, data := range databaseIzin {
					if data.Nomor == targetNomor {
						databaseIzin[index].Setujui()
						ditemukan = true
						break
					}
				}

				if !ditemukan {
					fmt.Println("❌ Nomor Izin tidak ditemukan di sistem.")
				} else {
					fmt.Println("✅ Izin berhasil disetujui!")
				}
			}

		case 4:
			fmt.Println("Menutup sistem...")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
