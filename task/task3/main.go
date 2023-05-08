package main

import "fmt"

func main() {
	var a int
	var c int

	fmt.Scan(&a)
	for i:=1; i <= a; i++ {
		c = i + 0 
		if i%3==0 || i%5 ==0 {
		fmt.Println(c) 	
		}
		
	} 
	
}
