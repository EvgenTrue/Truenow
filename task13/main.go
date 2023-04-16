package main

import "fmt"

func main() {

	var a int
	fmt.Println("введите число:")
	fmt.Scan(&a)

	if a <= 10 && a%10 == 0 {
		fmt.Fprintln(a / 10)
	}

}
