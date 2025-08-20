package konstanta

import "fmt"

func Pengandai(nama string, versi float64) {
	fmt.Printf("Halo, nama saya %s!\n", nama)
	fmt.Printf("Saya adalah aplikasi perkenalan sederhana yang dibuat dengan bahasa Go, versi %.1f.\n", versi)
	if nama == "GitHub Copilot" {
		fmt.Println("Saya adalah asisten AI yang siap membantu Anda dalam pengembangan perangkat lunak.")
	} else {
		fmt.Println("Saya adalah aplikasi perkenalan yang dibuat untuk tujuan pembelajaran.")
	}
	fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
}
