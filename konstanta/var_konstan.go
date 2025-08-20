package konstanta

import "fmt"

var x int = 10
var y int = 20
var z int = x + y

func Kalkulasi() {
	f := float32(x) / float32(y)
	fmt.Printf("Hasil penjumlahan %d + %d = %d\n", x, y, z)
	fmt.Printf("Hasil pembagian %d / %d = %.2f\n", x, y, f)
}
