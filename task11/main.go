package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("введите число")
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")

	}
	  { 
		{
			for i := 1; i <= b; i++ {
				for j := 1; j <= b-i; j-- {
					fmt.Print("*")
				}
				fmt.Println("")
			}
		}
	}
}
