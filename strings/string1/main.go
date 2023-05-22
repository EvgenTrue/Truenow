package main

import (
	"fmt"
)

func main() {
	// return true
	fmt.Println(isPalindrome("level"))

	// return false
	fmt.Println(isPalindrome("sample text"))

}
func isPalindrome(s string) bool {

	for i, j := 0, len(s)-1; i <= len(s)/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}

	}
	return true
}


