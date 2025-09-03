package array

import "fmt"

var months = [...]string{
	"Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

func ListMonths() {
	for i, month := range months {
		fmt.Printf("Bulan ke-%d adalah %s\n", i+1, month)
	}
}
