package main

import (
	"fmt"
	
)

func main() {

	a := "privet"
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(string(a[i]))
	}
	
}
