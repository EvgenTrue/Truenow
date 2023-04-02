package main

import "fmt"

func main() {
	var a = 10
	var c = 0

	fmt.Scan(&a)
	for i:=1; i <= a; i++ {
		c = i + c
		
		fmt.Println(c) 	
		}
		
	} 
	

