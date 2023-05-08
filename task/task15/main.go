package main

import "fmt"

func main() {
	var a int
	fmt.Print("введите число:")
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		for j := 1; j <= a; j++ {
			fmt.Print("*")
		}
		
		
		fmt.Println("")
		 
		

	}

}
