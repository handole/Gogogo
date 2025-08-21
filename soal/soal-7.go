package soal

import "fmt"

func Palindrome(kata string) {
	isPalindrome := true
	n := len(kata)

	for i := 0; i <= n/2; i++ {
		if kata[i] != kata[n-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println(kata, "Adalah Palindrome")
	} else {
		fmt.Println(kata, "Bukan Palindrome")
	}
}
