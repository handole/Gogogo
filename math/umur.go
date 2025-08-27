package math

import "fmt"

func Usia(umur int) {
	if umur < 0 {
		fmt.Println(("Umur tidak valid"))
	} else if umur <= 3 {
		fmt.Println("Bayi")
	} else if umur <= 5 {
		fmt.Println("Todler")
	} else if umur <= 12 {
		fmt.Println("Anak-anak")
	} else if umur <= 19 {
		fmt.Println("Remaja")
	} else if umur <= 59 {
		fmt.Println("Dewasa")
	} else {
		fmt.Println("Lansia")
	}
}
