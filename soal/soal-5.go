package soal

import "fmt"

func BigNumber(nums []int) {
	if len(nums) == 0 {
		fmt.Println("Array kosong.")
		return
	}

	biggest := nums[0]
	for _, num := range nums {
		if num > biggest {
			biggest = num
		}
	}

	fmt.Printf("Angka terbesar dalam array adalah: %d\n", biggest)
}
