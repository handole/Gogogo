package math

import "fmt"

func tambahan(a int, b int) int {
	return a + b
}

func kurangan(a int, b int) int {
	return a - b
}

func kalian(a int, b int) int {
	return a * b
}

func bagian(a int, b int) int {
	return a / b
}

func Kalkulator() {
	var a, b int
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&b)

	fmt.Printf("Hasil penjumlahan: %d\n", tambahan(a, b))
	fmt.Printf("Hasil pengurangan: %d\n", kurangan(a, b))
	fmt.Printf("Hasil perkalian: %d\n", kalian(a, b))
	fmt.Printf("Hasil pembagian: %d\n", bagian(a, b))
}
