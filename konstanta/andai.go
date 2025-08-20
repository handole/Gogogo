package konstanta

import "fmt"

// func Pengandai(nama string, versi float64) {
// 	fmt.Printf("Halo, nama saya %s!\n", nama)
// 	fmt.Printf("Saya adalah aplikasi perkenalan sederhana yang dibuat dengan bahasa Go, versi %.1f.\n", versi)
// 	if nama == "GitHub Copilot" {
// 		fmt.Println("Saya adalah asisten AI yang siap membantu Anda dalam pengembangan perangkat lunak.")
// 	} else {
// 		fmt.Println("Saya adalah aplikasi perkenalan yang dibuat untuk tujuan pembelajaran.")
// 	}
// 	fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
// }

func Pengandai() {
	nums := []int{1, 2, 3, 4, 5}
	for v := range nums {
		if v%2 == 0 {
			fmt.Printf("Angka %d adalah genap.\n", v)
		} else {
			fmt.Printf("Angka %d adalah ganjil.\n", v)
		}
	}
	fmt.Println("Pengandai selesai.")
}

func HariApa(day string) string {
	switch day {
	case "Senin":
		fmt.Println("Hari ini adalah Senin.")
		return "Hari Senin adalah hari pertama dalam seminggu."
	case "Selasa":
		fmt.Println("Hari ini adalah Selasa.")
		return "Hari Selasa adalah hari kedua dalam seminggu."
	case "Rabu":
		fmt.Println("Hari ini adalah Rabu.")
		return "Hari Rabu adalah hari ketiga dalam seminggu."
	case "Kamis":
		fmt.Println("Hari ini adalah Kamis.")
		return "Hari Kamis adalah hari keempat dalam seminggu."
	case "Jumat":
		fmt.Println("Hari ini adalah Jumat.")
		return "Hari Jumat adalah hari kelima dalam seminggu."
	case "Sabtu":
		fmt.Println("Hari ini adalah Sabtu.")
		return "Hari Sabtu adalah hari keenam dalam seminggu."
	case "Minggu":
		fmt.Println("Hari ini adalah Minggu.")
		return "Hari Minggu adalah hari ketujuh dalam seminggu."
	default:
		fmt.Println("Hari ini bukan hari yang valid.")
		return "Hari tidak valid."
	}
}
