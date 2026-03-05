package main

import "fmt"

// 1. KITA BUAT CETAKANNYA DULU
// TODO 1: Lengkapi Struct 'Barang' di bawah ini dengan 3 properti:
// - Nama (tipe: string)
// - Harga (tipe: int)
// - Stok (tipe: int)
type Barang struct {
	// ... (Tulis properti Anda di sini) ...
	Nama  string
	Harga int
	Stok  int
}

// 2. KITA BUAT METHOD DISKONNYA
// TODO 2: Buat fungsi bernama 'BeriDiskon' yang menempel pada pointer *Barang.
// Fungsi ini menerima satu parameter bernama 'persen' (tipe int).
// Rumus di dalamnya: b.Harga = b.Harga - ((b.Harga * persen) / 100)
// ... (Tulis kode method Anda di sini) ...

func (b *Barang) BeriDiskon(persen int) {
	b.Harga = b.Harga - ((b.Harga * persen) / 100)
}

func main() {
	// 3. SIAPKAN GUDANGNYA
	// TODO 3: Buat variabel bernama 'daftarBarang' yang merupakan Slice dari cetakan Barang.
	// ... (Tulis kode Anda di sini) ...
	var daftarBarang []Barang

	var pilihan int

	for {
		fmt.Println("\n=== 📦 Inventaris V3 (Struct & Pointer) ===")
		fmt.Println("1. Tambah Barang Baru")
		fmt.Println("2. Lihat Daftar Barang")
		fmt.Println("3. Berikan Diskon (Barang Pertama)")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu (1/2/3/4): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var inputNama string
			var inputHarga, inputStok int

			fmt.Print("Nama Barang (Tanpa Spasi): ")
			fmt.Scanln(&inputNama)
			fmt.Print("Harga Barang: ")
			fmt.Scanln(&inputHarga)
			fmt.Print("Sisa Stok: ")
			fmt.Scanln(&inputStok)

			// TODO 4: KEMAS DAN SIMPAN
			// a. Buat variabel baru bernama 'barangBaru' menggunakan cetakan Barang{ ... }
			//    Isi Nama, Harga, dan Stok dengan data dari input di atas.
			// b. Gunakan append() untuk memasukkan 'barangBaru' ke dalam slice 'daftarBarang'.
			// ... (Tulis kode Anda di sini) ...
			barangBaru := Barang{
				Nama:  inputNama,
				Harga: inputHarga,
				Stok:  inputStok,
			}

			daftarBarang = append(daftarBarang, barangBaru)

		case 2:
			fmt.Println("\n=== 📋 Daftar Barang ===")
			if len(daftarBarang) == 0 {
				fmt.Println("Gudang masih kosong!")
			} else {
				// TODO 5: TAMPILKAN BARANG
				// Lakukan for...range pada 'daftarBarang'.
				// Cetak urutan, Nama, Harga, dan Stok dari masing-masing barang.
				// ... (Tulis kode Anda di sini) ...
				for index, barang := range daftarBarang {
					fmt.Printf("%d. %s - Rp %d (Sisa Stok: %d unit)\n", index+1, barang.Nama, barang.Harga, barang.Stok)
				}
			}

		case 3:
			// Skenario Sederhana: Kita berikan diskon khusus untuk barang yang paling pertama masuk (index 0)
			if len(daftarBarang) == 0 {
				fmt.Println("❌ Gudang masih kosong, tidak ada yang bisa didiskon!")
			} else {
				var persenDiskon int
				fmt.Printf("Masukkan persentase diskon untuk %s: ", daftarBarang[0].Nama)
				fmt.Scanln(&persenDiskon)

				// TODO 6: EKSEKUSI METHOD
				// Panggil method BeriDiskon() pada daftarBarang[0] dan masukkan nilai 'persenDiskon' ke dalam kurungnya.
				// ... (Tulis kode Anda di sini) ...

				daftarBarang[0].BeriDiskon(persenDiskon)
				fmt.Println("✅ Diskon berhasil diterapkan! Silakan cek menu nomor 2.")
			}

		case 4:
			fmt.Println("Keluar dari aplikasi...")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
