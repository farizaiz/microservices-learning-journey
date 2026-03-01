package main

import "fmt"

func main() {
	fmt.Println("=== Kalkulator Sederhana ====")

	var angka1, angka2 float64
	var operator string
	var hasil float64

	fmt.Print("Masukkan Angka Pertama: ")
	fmt.Scanln(&angka1)

	fmt.Print("Masukkan Angka Kedua: ")
	fmt.Scanln(&angka2)

	fmt.Print("Masukkan Operator: ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		hasil = angka1 + angka2
	case "-":
		hasil = angka1 - angka2
	case "*":
		hasil = angka1 * angka2
	case "/":
		if angka2 == 0 {
			fmt.Println("Error, tidak bisa membagi dengan 0")
			return
		}
	default:
		fmt.Println("Operator Tidak Dikenal")
		return
	}

	fmt.Println("Hasil Perhitungan ", angka1, operator, angka2, "Adalah: ", hasil)

}
