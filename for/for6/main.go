package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("vvedi chislo:")
	fmt.Scan(&a)
	fmt.Print("vvedi chislo:")
	fmt.Scan(&b)
	for i := b; i >= a; i-- {

		fmt.Println(i)
	}
}
