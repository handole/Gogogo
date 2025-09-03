package array

import "fmt"

func ListNumbers() {
	numbers := []int{32, 12, 54, 76, 223, 143, 65, 169}

	// fmt.Println("Isi array numbers:", numbers[4])
	numbers = append(numbers, 99, 100, 101)

	for index, value := range numbers {
		fmt.Printf("Index ke-%d : %d\n", index, value)
	}
}

func ListPop(slice []int) ([]int, int) {
	if len(slice) == 0 {
		return slice, 0 // Return the original slice and a zero value if it's empty
	}
	val := slice[len(slice)-1]
	return slice[:len(slice)-1], val

}

func ListCapacity() {
	numbers := make([]int, 5, 10) // Slice dengan panjang 5 dan kapasitas 10
	fmt.Printf("Panjang: %d, Kapasitas: %d\n", len(numbers), cap(numbers))
}
