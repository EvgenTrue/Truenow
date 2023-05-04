package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"man", "jam", "apple", "rave", "cacao"}
	for i := 0; i < len(a); i++ {
		b := strings.Replace(a[i], "a", "A", 2)
		fmt.Println(b)
	}
}
