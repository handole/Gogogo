package konstanta

import (
	"encoding/json"
	"fmt"
	"os"
)

type Child struct {
	Title       string `json:"title"`
	Method      string `json:"method"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"isDefault"`
	Logo        string `json:"logo"`
}

type MethodPayment struct {
	PaymentGroupTitle string  `json:"paymentGroupTitle"`
	Childs            []Child `json:"childs"`
	Code              string  `json:"code"`
}

func LoadPaymentMethods() {
	// Baca file JSON
	file, err := os.ReadFile("data/payment_method.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// parsing json to struct
	var methods []MethodPayment
	err = json.Unmarshal(file, &methods)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Tampilkan hasil
	for _, method := range methods {
		fmt.Println("Payment Group Title:", method.PaymentGroupTitle)
		// for _, child := range method.Childs {
		// 	fmt.Println("  Title:", child.Title)
		// 	fmt.Println("  Method:", child.Method)
		// 	fmt.Println("  Code:", child.Code)
		// 	fmt.Println("  Name:", child.Name)
		// 	fmt.Println("  Description:", child.Description)
		// 	fmt.Println("  Is Default:", child.IsDefault)
		// 	fmt.Println("  Logo:", child.Logo)
		// }
		fmt.Println("Code:", method.Code)
	}
}
