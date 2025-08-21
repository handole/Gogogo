package soal

import "fmt"

func TebakHari(day string) {
	switch day {
	case "Sabtu", "Minggu":
		fmt.Println("Hari ini adalah Weekend.")
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari ini adalah Weekday.")
	default:
		fmt.Println("Hari ini bukan hari yang valid.")
	}
}
