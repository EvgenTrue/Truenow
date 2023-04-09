package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("vvedi 2 chisla")
	fmt.Scan(&a, &b)
	if a < b {
		fmt.Println(a, b)
	}
	if a > b {
		fmt.Println(b, a)
	}
	
	
}
