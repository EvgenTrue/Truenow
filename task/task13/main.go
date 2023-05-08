package main

import "fmt"

func main() {

	var a int
	fmt.Println("введите число:")
	fmt.Scan(&a)

	for a > 0 {
		fmt.Print(a % 10)
		a /= 10

	}

}

/// 1546
/// %10 = 6 
/// 154 