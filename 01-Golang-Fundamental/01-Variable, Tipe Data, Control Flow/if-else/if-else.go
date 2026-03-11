package main

import "fmt"

func main() {

	pilihan := 1

	if pilihan == 1 {
		fmt.Println("Anda Memilih Penjumlahan")
	} else if pilihan == 2 {
		fmt.Println("Anda Memilih Pengurangan")
	} else {
		fmt.Println("Pilihan Tidak Valid")
	}
}
