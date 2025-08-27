package math

import "fmt"

func TebakAngka(tebakan int) (bool, string) {
	angkaRahasia := 7
	if tebakan < 1 || tebakan > 10 {
		return false, "Tebakan harus antara 1 hingga 10."
	}
	if tebakan < angkaRahasia {
		return false, "Tebakan terlalu rendah."
	}
	if tebakan > angkaRahasia {
		return false, "Tebakan terlalu tinggi."
	}
	return true, "Selamat! Tebakan Anda benar."
}

func Tebakanya() {
	var tebakan int

	fmt.Println("Tebak angka antara 1 - 10")
	fmt.Print("Masukkan tebakanmu: ")
	fmt.Scan(&tebakan)

	// langsung dapet string
	benar, pesan := TebakAngka(tebakan)
	fmt.Println(pesan)

	if benar {
		fmt.Println("Game selesai ✅")
	} else {
		fmt.Println("Coba lagi ❌")

	}
}
