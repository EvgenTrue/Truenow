package main

import "fmt"

func main() {

	var a, b int
	b = a / 2
	fmt.Print("введите число:")
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		for j := 1; j <= b; i++ {
			fmt.Print("")
		}
	}
}
