package main

import "fmt"

func main() {
	var a, c int
	fmt.Scanln(&a)
	for i:=1; i <= a; i++ {
		if c%3 || c%5 {
			fmt.Println(c) 
		}
	}
	} 