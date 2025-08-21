package soal

import "fmt"

func GanjilGenap(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Printf("Angka %d adalah Genap.\n", i)
		} else {
			fmt.Printf("Angka %d adalah Ganjil.\n", i)
		}
	}
}
