package main

import (
	"fmt"
)

func main() {
	a := "good day"

	s := []string{"a", "o", "e", "i", "u"}
	for i := 0; i < len(a); i++ {
		for _, j := range s {
			if string(a[i]) == j {
				
				fmt.Println(i, string(a[i]))
				break
				
			}
		
		}

	}
}
