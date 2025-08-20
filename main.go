package main

import (
	"Gogogo/konstanta"
	"fmt"
)

func main() {
	konstanta.Perkenalan()
	fmt.Println("================================================")
	konstanta.PrintKonstanta()
	fmt.Println("================================================")
	konstanta.Kalkulasi()
	fmt.Println("================================================")
	konstanta.HariApa("Rabu")
	fmt.Println("================================================")
	konstanta.LoadPaymentMethods()
	fmt.Println("================================================")
	fmt.Println("Program selesai.")
}
