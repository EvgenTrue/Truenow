package main

import "fmt"

func main() {
	var a, b int
	
	fmt.Println("введите число")
	fmt.Scan(&a)
	for i := 0; i <= a; i++ {
		
		if i<=a/2 {
			b++
		}else {
			b--
		}
		
		for j := 1; j <= b; j++ {
			fmt.Print("*")
		}
		fmt.Println("")

	}
}
	

